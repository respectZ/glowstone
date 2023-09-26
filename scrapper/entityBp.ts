import {
  DOMParser,
  Element,
} from "https://deno.land/x/deno_dom@v0.1.38/deno-dom-wasm.ts";

const URL = "https://bedrock.dev/docs/stable/Entities";
// const RESULT: string[] = [];
const RESULT: Record<string, string[]> = {};

const PascalToSnake = (str: string) =>
  str.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`).slice(1);

const snakeToPascal = (str: string) =>
  str
    .split("/")
    .map((snake) =>
      snake
        .split("_")
        .map((s) =>
          s
            .split(".")
            .map((substr) => substr.charAt(0).toUpperCase() + substr.slice(1))
            .join("_")
        )
        .join("")
    )
    .join("/");
const write = (dest: string, data: string) =>
  Deno.writeTextFileSync(dest, data);

interface GoField {
  name: string;
  type: string;
  description?: string;
}

function toGoStruct(
  name: string,
  description: string,
  fields: GoField[]
): string {
  // Change name from minecraft:*.* to *.*
  name = name.replace("minecraft:", "");
  name = snakeToPascal(name);
  // Check if component is a behavior. (Behavior_). If so, add propty to the field
  if (name.includes("Behavior_")) {
    const priority = <GoField>{
      name: "priority",
      type: "int",
      description:
        "The priority of this behavior. Lower values are higher priority.",
    };
    fields.push(priority);
  }

  const temp: string[] = []; // Track field name to prevent duplicates.

  let result = `// ${description}\n`;
  result += `type ${name} struct {\n`;
  for (const [_, value] of Object.entries(fields)) {
    const fieldName = snakeToPascal(value.name);
    if (temp.includes(fieldName)) continue;
    temp.push(fieldName);

    result += `  // ${value.description}\n`;
    result += `  ${fieldName} ${value.type} \`json:"${value.name},omitempty"\`\n`;
  }
  result += "}\n";
  return result;
}

function newGoStruct(name: string, description: string, fields: GoField[]) {
  let struct = "";
  struct += `// ${description}\n`;
  struct += `type ${snakeToPascal(name)} struct {\n`;
  for (let i = 0; i < fields.length; i++) {
    const field = fields[i];
    struct += `    // ${field.description}\n`;
    struct += `    ${field.name} ${field.type} \`json:"${PascalToSnake(
      field.name
    )},omitempty"\`\n\n`;
  }
  struct += "}";
  return struct;
}

function getFields(table: Element, parentName: string): GoField[] {
  const fields: GoField[] = [];
  const rows = [...table.children[0].children];
  for (let i = 1; i < rows.length; i++) {
    const row = rows[i];
    const field: GoField = {
      name: row.children[0].innerText,
      type: row.children[1].innerText,
      description: row.children[3].innerText,
    };
    // Change type to go type
    switch (field.type) {
      // Primitive types
      case "String":
        field.type = "string";
        break;
      case "Number":
        field.type = "int";
        break;
      case "Boolean":
        field.type = "bool";
        break;
      case "Integer":
        field.type = "int";
        break;
      case "Decimal":
        field.type = "float64";
        break;
      case "Positive Integer":
        field.type = "uint";
        break;
      // Non primitive types
      case "Item Description Properties":
        field.type = "string";
        break;
      case "List":
      case "Array":
        field.type = "[]interface{}";
        if (field.description)
          field.description = field.description.split("\n")[0];
        break;
      case "Trigger":
        field.type = "interface{}";
        field.description = "[Trigger] " + field.description;
        break;
      case "Range [a, b]": // TODO: check int or float
        field.type = "float64";
        break;
      case "Minecraft Filter":
        field.type = "*f.Filter";
        break;
      case "Vector [a, b, c]":
        field.type = "[]float64";
        field.description = "[Vector3 [a,b,c]] " + field.description;
        break;
      case "Molang":
        field.type = "string";
        field.description = "[Molang String] " + field.description;
        break;
      case "JSON Object":
        field.type = "interface{}";
        if (field.description)
          field.description = field.description.split("\n")[0];
        break;
    }
    // Need to check if description children+1 is table, so it's JSON Object, and then convert it to go struct
    if (row.children[3].children.length > 0) {
      for (let j = 0; j < row.children[3].children.length; j++) {
        const child = row.children[3].children[j];
        if (child.tagName == "TABLE") {
          // Change type to struct
          let fieldName = parentName.replace("minecraft:", "");
          fieldName = fieldName.replace(/\./g, "_");
          fieldName = snakeToPascal(fieldName);

          const fieldType = fieldName + snakeToPascal(field.name);

          field.description = child.parentElement?.innerText.split("\n")[0];

          // Convert table to go struct
          const childFields = getFields(child, parentName);
          const childStruct = toGoStruct(
            fieldType,
            "[Not a component] " + field.description,
            childFields
          );
          //   RESULT.push(childStruct);
          if (!RESULT[parentName]) RESULT[parentName] = [];
          RESULT[parentName] = [...RESULT[parentName], childStruct];

          // Check if field type is "List"
          if (field.type.includes("[]interface{}")) {
            field.type = "[]" + fieldType;
          } else {
            // Need to check if the fieldName is EntityTypes => interface{}, since it can be Array or struct, like in nearest attackable target.
            if (field.name == "entity_types") {
              field.description = `Corresponding Type: ${fieldType}\n`;
            } else {
              field.type = fieldType;
            }
          }
        }
      }
    }
    fields.push(field);
  }
  return fields;
}

async function main() {
  const document = new DOMParser().parseFromString(
    await (await fetch(URL)).text(),
    "text/html"
  );
  if (!document) throw new Error("Failed to parse document");

  const components = Array.from(
    document.querySelectorAll("h2.anchored-heading")
  )
    .map((c) => {
      const component = <Element>c;
      return component;
    })
    .filter((c) => c.innerText.includes("minecraft:"));
  const structs: string[] = [];

  components.forEach((component) => {
    const name = component.innerText.replace("#", "");
    let description = component.nextElementSibling?.innerText;
    let table = component.nextElementSibling?.nextElementSibling;
    let fields: GoField[] = [];

    // Check if description is table
    if (component.nextElementSibling?.tagName == "TABLE") {
      table = component.nextElementSibling;
      description = "\n";
    }

    // Check for duplicate
    if (structs.includes(name)) return;

    // Check if next element after sibling is table
    if (table?.tagName == "TABLE") {
      fields = getFields(table, name);
    }

    if (!description) description = "";

    // Remove newlines
    description = description.replace(/\n/g, " ");

    structs.push(name);
    // RESULT.push(toGoStruct(name, description, fields));
    if (!RESULT[name]) RESULT[name] = [];
    RESULT[name] = [...RESULT[name], toGoStruct(name, description, fields)];
  });

  //   Add custom mstructs
}

/******* Custom Components, either unable to do automatically or missing. */
function health() {
  const field: GoField[] = [
    {
      name: "Value",
      type: "int",
      description: "The amount of health the entity has.",
    },
    {
      name: "Max",
      type: "int",
      description: "The maximum amount of health the entity can have.",
    },
    {
      name: "Min",
      type: "int",
      description: "The minimum amount of health the entity can have.",
    },
  ];
  return newGoStruct("Health", "Defines the health of this entity.", field);
}
function ambientSoundInterval_EventName() {
  const fields: GoField[] = [
    {
      name: "Condition",
      type: "string",
      description:
        "List of dynamic level sound events, with conditions for choosing between them. Evaluated in order, first one wins. If none evaluate to true, 'event_name' will take precedence.",
    },
    {
      name: "EventName",
      type: "string",
      description: "Level sound event to be played as the ambient sound",
    },
  ];
  return newGoStruct(
    "AmbientSoundInterval_EventNames",
    "Defines the sound events that can be played by this entity.",
    fields
  );
}

function angerLevel_OnIncreaseSounds() {
  const fields: GoField[] = [
    {
      name: "Condition",
      type: "string",
      description:
        "A Molang expression describing under which conditions to play this sound, given that the entity was provoked",
    },
    {
      name: "Sound",
      type: "string",
      description: "The sound to play",
    },
  ];
  return newGoStruct(
    "AngerLevel_OnIncreaseSounds",
    "Sounds to play when the entity is getting provoked. Evaluated in order. First matching condition wins.",
    fields
  );
}

function hurtOnCondition_DamageConditions() {
  const fields: GoField[] = [
    {
      name: "Cause",
      type: "string",
      description:
        "The kind of damage that is caused to the entity. Various armors and spells use this to determine if the entity is immune.",
    },
    {
      name: "DamagePerTick",
      type: "int",
      description:
        "The amount of damage done each tick that the conditions are met.",
    },
    {
      name: "Filters",
      type: "*f.Filter",
      description:
        "The set of conditions that must be satisfied before the entity takes the defined damage.",
    },
  ];
  return newGoStruct(
    "HurtOnCondition_DamageConditions",
    "List of damage conditions that when met can cause damage to the entity.",
    fields
  );
}

function shooter_Projectile() {
  const fields: GoField[] = [
    {
      name: "Def",
      type: "string",
      description:
        "The identifier of the projectile definition to use for this projectile.",
    },
    {
      name: "AuxVal",
      type: "int",
      description:
        "The aux value of the projectile. This is used to determine the type of projectile to spawn.",
    },
    {
      name: "Power",
      type: "float64",
      description:
        "The power of the projectile. This is used to determine the velocity of the projectile.",
    },
    {
      name: "Magic",
      type: "bool",
      description:
        "Sets whether the projectiles being used are flagged as magic. If set, the ranged attack goal will not be used at the same time as other magic goals, such as minecraft:behavior.drink_potion",
    },
    {
      name: "Sound",
      type: "string",
      description: "Sound that is played when the shooter shoots a projectile.",
    },
    {
      name: "Chance",
      type: "float64",
      description:
        "The chance that this projectile will be fired when the shooter shoots.",
    },
    {
      name: "LoseTarget",
      type: "bool",
      description: "When set to true, the entity will shoot this projectile.",
    },
    {
      name: "Filters",
      type: "*f.Filter",
      description:
        "The set of conditions that must be satisfied before the entity shoots this projectile.",
    },
  ];
  return newGoStruct(
    "Shooter_Projectile",
    "List of projectiles that can be used by the shooter. Projectiles are evaluated in the order of the list; After a projectile is chosen, the rest of the list is ignored.",
    fields
  );
}

await main();

// Push custom components
RESULT["minecraft:health"] = [health()];
RESULT["minecraft:ambient_sound_interval"] = [
  ...RESULT["minecraft:ambient_sound_interval"],
  ambientSoundInterval_EventName(),
];
RESULT["minecraft:anger_level"] = [
  ...RESULT["minecraft:anger_level"],
  angerLevel_OnIncreaseSounds(),
];
RESULT["minecraft:hurt_on_condition"] = [
  ...RESULT["minecraft:hurt_on_condition"],
  hurtOnCondition_DamageConditions(),
];
RESULT["minecraft:shooter"] = [
  ...RESULT["minecraft:shooter"],
  shooter_Projectile(),
];

// write("./test.go", RESULT.join("\n"));
for (const [key, value] of Object.entries(RESULT)) {
  const name = key.split(":")[1];
  let data = "package component\n\n";
  data += "// Code generated by scrapper/entityBp.ts;\n\n";
  const strct = value.join("\n");
  const imports: Record<string, string> = {};

  //   Add filter
  if (strct.includes("*f.Filter")) {
    imports["f"] = `github.com/respectZ/glowstone/bp/types`;
  }

  //   Add imports
  if (Object.keys(imports).length > 0) {
    data += "import (\n";
    for (const [key, value] of Object.entries(imports)) {
      data += `  ${key} "${value}"\n`;
    }
    data += ")\n\n";
  }

  data += strct;

  write(`./bp/entity/component/${name}.go`, data);
}
