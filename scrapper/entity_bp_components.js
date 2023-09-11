let components = [...document.querySelectorAll("h2.anchored-heading")].filter(
  (e) => e.innerText.indexOf("minecraft:") != -1
);
let parent = [...document.querySelectorAll("div.docs-content")[0].children];
let result = [];

function camelToPascal(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

const snakeToPascal = (string) => {
  return string
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
};

/**
 * @param {string} name
 * @param {fields} fields
 * @returns {string}
 * @example
 * newGoStruct("minecraft:ambient_sound_interval", [{
 *    "name": "condition",
 *    "type": "String",
 *    "description": "A Molang expression describing under which conditions to play this sound, given that the entity was provoked",
 * }])
 */
function newGoStruct(name, description, fields) {
  let struct = "";
  struct += `// ${description}\n`;
  struct += `type ${snakeToPascal(name)} struct {\n`;
  for (var i = 0; i < fields.length; i++) {
    let field = fields[i];
    struct += `    // ${field.description}\n`;
    struct += `    ${field.name} ${field.type} \`json:"${field.name},omitempty"\`\n\n`;
  }
  struct += "}";
  return struct;
}

/*** Special cases for Array type */
function ambientSoundInterval_EventNames() {
  let struct = "";
  struct +=
    "// List of dynamic level sound events, with conditions for choosing between them. Evaluated in order, first one wins. If none evaluate to true, 'event_name' will take precedence.\n";
  struct += "type AmbientSoundInterval_EventNames struct {\n";
  struct +=
    "    // List of dynamic level sound events, with conditions for choosing between them. Evaluated in order, first one wins. If none evaluate to true, 'event_name' will take precedence.\n";
  struct += '    Condition string `json:"condition,omitempty"`\n\n';
  struct += "    // Level sound event to be played as the ambient sound\n";
  struct += '    EventName string `json:"event_name,omitempty"`\n\n';
  struct += "}";
  return struct;
}

function angerLevel_OnIncreaseSounds() {
  let struct = "";
  struct +=
    "// Sounds to play when the entity is getting provoked. Evaluated in order. First matching condition wins\n";
  struct += "type AngerLevelOn_IncreaseSounds struct {\n";
  struct +=
    "    // A Molang expression describing under which conditions to play this sound, given that the entity was provoked\n";
  struct += '    Condition string `json:"condition,omitempty"`\n\n';
  struct += "    // The sound to play\n";
  struct += '    Sound string `json:"sound,omitempty"`\n\n';
  struct += "}";
  return struct;
}

function hurtOnCondition_DamageConditions() {
  let struct = "";
  struct +=
    "// List of damage conditions that when met can cause damage to the entity.\n";
  struct += "type HurtOnCondition_DamageConditions struct {\n";
  struct +=
    "    // The kind of damage that is caused to the entity. Various armors and spells use this to determine if the entity is immune.\n";
  struct += '    Cause string `json:"cause,omitempty"`\n\n';
  struct +=
    "    // The amount of damage done each tick that the conditions are met.\n";
  struct += '    DamagePerTick int `json:"damage_per_tick,omitempty"`\n\n';
  struct +=
    "    // The set of conditions that must be satisfied before the entity takes the defined damage.\n";
  struct += '    Filters *f.Filter `json:"filters,omitempty"`\n\n';
  struct += "}";
  return struct;
}

/*** End of special cases for Array type */

/*** Special cases for List type */

function behaviorDrinkPotion_Potions() {
  let fields = [
    {
      name: "chance",
      type: "float64",
      description:
        "The percent chance (from 0.0 to 1.0) of this potion being selected when searching for a potion to use.",
    },
    {
      name: "filters",
      type: "*f.Filter",
      description:
        "The filters to use when determining if this potion can be selected.",
    },
    {
      name: "id",
      type: "string",
      description: "The registry ID of the potion to use",
    },
  ];
  return newGoStruct(
    "BehaviorDrinkPotion_Potions",
    "A list of potions that this entity can drink.",
    fields
  );
}

function behaviorSendEvent_Sequence() {
  let fields = [
    {
      name: "base_delay",
      type: "float64",
      description: "Amount of time in seconds before starting this step",
    },
    {
      name: "event",
      type: "string",
      description: "The event to send to the entity",
    },
    {
      name: "sound_event",
      type: "string",
      description: "The sound event to play when this step happens",
    },
  ];
  return newGoStruct(
    "BehaviorSendEvent_Sequence",
    "List of events to send",
    fields
  );
}

function conditionalBandwithOptimization_ConditionalValues() {
  let fields = [
    {
      name: "Conditional_values",
      type: "*f.Filter",
      description:
        "Conditions that must be met for these optimization values to be used.",
    },
    {
      name: "Max_dropped_ticks",
      type: "int",
      description:
        "In relation to the optimization value, determines the maximum ticks spatial update packets can be not sent.",
    },
    {
      name: "Max_optimized_distance",
      type: "float64",
      description:
        "The maximum distance considered during bandwidth optimizations. Any value below the max is interpolated to find optimization, and any value greater than or equal to this max results in max optimization.",
    },
    {
      name: "Use_motion_prediction_hints",
      type: "bool",
      description:
        "When set to true, smaller motion packets will be sent during drop packet intervals, resulting in the same amount of packets being sent as without optimizations but with much less data being sent. This should be used when actors are travelling very quickly or teleporting to prevent visual oddities.",
    },
  ];
  return newGoStruct(
    "ConditionalBandwithOptimization_ConditionalValues",
    "Defines the Conditional Spatial Update Bandwidth Optimizations of this entity.",
    fields
  );
}

/*** End of special cases for List type */

/*** Special cases for JSON Object type */

function conditionalBandwithOptimization_DefaultValues() {
  let fields = [
    {
      name: "Max_dropped_ticks",
      type: "int",
      description:
        "In relation to the optimization value, determines the maximum ticks spatial update packets can be not sent.",
    },
    {
      name: "Max_optimized_distance",
      type: "float64",
      description:
        "The maximum distance considered during bandwidth optimizations. Any value below the max is interpolated to find optimization, and any value greater than or equal to this max results in max optimization.",
    },
    {
      name: "Use_motion_prediction_hints",
      type: "bool",
      description:
        "When set to true, smaller motion packets will be sent during drop packet intervals, resulting in the same amount of packets being sent as without optimizations but with much less data being sent. This should be used when actors are travelling very quickly or teleporting to prevent visual oddities.",
    },
  ];
  return newGoStruct(
    "ConditionalBandwithOptimization_DefaultValues",
    "Defines the Conditional Spatial Update Bandwidth Optimizations of this entity.",
    fields
  );
}

function health() {
  let fields = [
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
  return newGoStruct("Health", "Defines the health of this entity.", fields);
}

/*** End of cases for JSON Object type */

/*** Special cases for st*/

function getFields(table, parentName) {
  let fields = [];
  let rows = [...table.children[0].children];
  for (var i = 1; i < rows.length; i++) {
    let row = rows[i];
    let field = {
      name: row.children[0].innerText,
      type: row.children[1].innerText,
      defaultVal: row.children[2].innerText,
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
        field.type = "[]interface{}";
        // console.log(parentName);
        field.description = field.description.split("\n")[0] + "\n";
        break;
      case "Array":
        // if(parentName == "minecraft:ambient_sound_interval") {
        //     field.type = "[]AmbientSoundIntervalEventNames";
        // }
        // if(parentName == "minecraft:anger_level") {
        //     field.type = "[]AngerLevelOnIncreaseSounds";
        // }
        // if(parentName == "minecraft:healable") {
        //     field.type = "[]HealableItems";
        // }
        // if(parentName == "minecraft:hurt_on_condition") {
        //     field.type = "[]HurtOnConditionDamageConditions";
        // }
        field.type = "[]interface{}";
        field.description = field.description.split("\n")[0] + "\n";
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
        field.description = field.description.split("\n")[0] + "\n";
        break;
    }
    // Need to check if description children+1 is table, so it's JSON Object, and then convert it to go struct
    if (row.children[3].children.length > 0) {
      for (var j = 0; j < row.children[3].children.length; j++) {
        let child = row.children[3].children[j];
        if (child.tagName == "TABLE") {
          // Change type to struct
          let fieldName = parentName.replace("minecraft:", "");
          fieldName = fieldName.replace(/\./g, "_");
          fieldName = snakeToPascal(fieldName);

          let fieldType = fieldName + snakeToPascal(field.name);

          field.description =
            child.parentElement.innerText.split("\n")[0] + "\n";

          // Convert table to go struct
          let childFields = getFields(child, parentName);
          let childStruct = toGoStruct(
            fieldType,
            "[Not a component] " + field.description,
            childFields
          );
          result.push(childStruct);

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

function toGoStruct(name, description, fields) {
  // Change name from minecraft:*.* to *.*
  name = name.replace("minecraft:", "");
  // name = name.replace(/\./g, "_");
  name = snakeToPascal(name);
  // Check if component is a behavior. (Behavior_). If so, add propty to the field
  if (name.includes("Behavior_")) {
    let f = {
      name: "priority",
      type: "int",
      defaultVal: "0",
      description:
        "The priority of this behavior. Lower values are higher priority.\n",
    };
    fields.push(f);
  }
  let temp = []; // Track field name to prevent duplicates.

  let result = "// " + description;
  result += "type " + name + " struct {\n";
  for (var i = 0; i < fields.length; i++) {
    let field = fields[i];
    let fieldName = snakeToPascal(field.name);
    if (temp.includes(fieldName)) continue;
    temp.push(fieldName);
    result += "    // " + field.description;
    result +=
      "    " +
      fieldName +
      " " +
      field.type +
      ' `json:"' +
      field.name +
      ',omitempty"`\n\n';
  }
  result += "}\n";
  return result;
}

function main() {
  let structs = []; // track structs to avoid duplicates
  for (var i = 0; i < components.length; i++) {
    let component = components[i];
    let name = component.innerText;
    let description = component.nextElementSibling.innerText;
    let table = component.nextElementSibling.nextElementSibling;
    let fields = [];
    // Check if description is table
    if (component.nextElementSibling.tagName == "TABLE") {
      table = component.nextElementSibling;
      description = "\n";
    }

    // Check for duplicate structs
    if (structs.includes(name)) {
      continue;
    }
    // Check if next element after sibling is table
    if (table.tagName == "TABLE") {
      fields = getFields(table, name);
    }
    structs.push(name);
    result.push(toGoStruct(name, description, fields));
  }
  // Add custom structs
  result.push(angerLevel_OnIncreaseSounds());
  result.push(ambientSoundInterval_EventNames());
  result.push(hurtOnCondition_DamageConditions());

  result.push(conditionalBandwithOptimization_ConditionalValues());
  result.push(conditionalBandwithOptimization_DefaultValues());
  result.push(health());
}

main();

console.log(result.join("\n"));

// TODO:
