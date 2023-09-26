// https://learn.microsoft.com/en-us/minecraft/creator/reference/content/itemreference/examples/itemcomponentlist

import {
  DOMParser,
  Element,
} from "https://deno.land/x/deno_dom@v0.1.38/deno-dom-wasm.ts";

interface BaseItemComponent {
  name: string;
  url: string;
  description: string;
}

interface GoField {
  name: string;
  type: string;
  description?: string;
}

interface GoStruct {
  name: string;
  fields: GoField[];
  description?: string;
}

const GoDataMap: { [key: string]: string } = {
  string: "string",
  int: "int",
  integer: "int",
  float: "float64",
  boolean: "bool",
  array: "[]interface{}",
  list: "[]interface{}",
  "vector [x,y,z]": "[]float64",
  event: "string",
  "integer range": "[]int",
  trigger: "*types.Trigger",
  "json object": "interface{}",
  enumerator: "string", // minecraft:wearable
  "": "int", // minecraft:max_stack_size
  "--": "interface{}", // minecraft:enchantable
};

function BaseItemComponentFromTR(tr: Element): BaseItemComponent {
  const name = tr.querySelector("td:nth-child(1)")?.textContent?.trim() ?? "";
  const url = tr.querySelector("td:nth-child(1) a")?.getAttribute("href") ?? "";
  const description =
    tr.querySelector("td:nth-child(4)")?.textContent?.trim() ?? "";

  return {
    name,
    url,
    description,
  };
}

async function GoStructFromURL(url: string): Promise<GoStruct[]> {
  // Fetch
  const res = await fetch(url);
  const html = await res.text();

  const document = new DOMParser().parseFromString(html, "text/html");
  if (!document) throw new Error("Failed to parse HTML");

  // Find tables
  const _tables = document.querySelectorAll("table");

  const r: GoStruct[] = [];

  _tables.forEach((_table) => {
    // Find description, which is n-1 of table
    const table = <Element>_table;
    const el = table.previousElementSibling;
    // If el is p, then it's description
    // If el is h2, then it's name
    if (!el) throw new Error("Failed to find description");

    let description = "";
    let name = "";
    switch (el.tagName) {
      case "P":
        {
          description = el.textContent?.trim() ?? "";
          // Get Name
          const _name = el.previousElementSibling;
          if (!_name) throw new Error("Failed to find name");
          if (_name.tagName !== "H2") throw new Error("Failed to find name");
          name = _name.textContent?.trim() ?? "";
        }
        break;
      case "H2":
        name = el.textContent?.trim() ?? "";
        break;
    }

    // Get fields
    const fields = TableToFields(table);
    const struct: GoStruct = {
      name,
      fields,
      description,
    };
    r.push(struct);
  });
  return r;
}

function TableToFields(table: Element): GoField[] {
  // Loop tbody and tr
  const _tbody = table.querySelector("tbody");
  if (!_tbody) throw new Error("Failed to find tbody");

  // Loop through tr
  const _trs = _tbody.querySelectorAll("tr");
  const r: GoField[] = [];
  _trs.forEach((_tr) => {
    const _t = <Element>_tr;
    const name = _t.querySelector("td:nth-child(1)")?.textContent?.trim() ?? "";
    const _type =
      _t.querySelector("td:nth-child(3)")?.textContent?.trim() ?? "";

    // Check type from datamap
    const gtype = GoDataMap[_type.toLowerCase()] ?? _type;
    if (GoDataMap[_type.toLowerCase()] === undefined)
      console.log(`Unknown type: ${_type} in ${name}`);

    const description =
      _t.querySelector("td:nth-child(4)")?.textContent?.trim() ?? "";

    r.push(<GoField>{
      name,
      type: gtype,
      description,
    });
  });
  return r;
}

function InterfaceToGoStruct(data: GoStruct): string {
  let r = ``;
  r += `// ${data.description}\n`;
  r += `type ${data.name} struct {\n`;
  data.fields.forEach((f) => {
    const fieldName = lowerCaseToCamelCase(f.name);
    r += `  // ${f.description}\n`;
    r += `  ${fieldName} ${f.type} \`json:"${f.name},omitempty"\`\n`;
    r += "\n";
  });
  r += "}\n";
  return r;
}

function write(dest: string, data: string): void {
  Deno.writeTextFileSync(dest, data);
}

function lowerCaseToCamelCase(str: string): string {
  let s = str.replace(/_([a-z])/g, function (g) {
    return g[1].toUpperCase();
  });
  s = s.charAt(0).toUpperCase() + s.slice(1);
  return s;
}

(async () => {
  const url =
    "https://learn.microsoft.com/en-us/minecraft/creator/reference/content/itemreference/examples/itemcomponentlist";

  // Fetch
  const res = await fetch(url);
  const html = await res.text();

  const document = new DOMParser().parseFromString(html, "text/html");
  if (!document) throw new Error("Failed to parse HTML");

  const _table = document.querySelector("table");
  if (!_table) throw new Error("Failed to find table");

  //   Find tbody
  const _tbody = _table.querySelector("tbody");
  if (!_tbody) throw new Error("Failed to find tbody");

  // Loop through tr
  const _trs = _tbody.querySelectorAll("tr");
  for (let i = 0; i < _trs.length; i++) {
    const _tr = _trs[i];
    const _t = <Element>_tr;
    const r = BaseItemComponentFromTR(_t);
    if (
      r.name === "description" ||
      r.name === "menu_category" ||
      r.name === "item"
    )
      continue;

    const ps = await GoStructFromURL(
      `https://learn.microsoft.com/en-us/minecraft/creator/reference/content/itemreference/examples/${r.url}`
    );

    let structs: string[] = [
      "package component\n",
      "// Code generated by scrapper/item/main.ts;",
      "",
    ];
    // To struct go
    let index = 0;
    ps.forEach((p) => {
      // Replace p.name if it's Parameter
      if (p.name === "Parameters") {
        p.name = lowerCaseToCamelCase(r.name.split(":")[1]);
        if (p.description === "") p.description = r.description;
      }

      //   If p index is > 0, then add base name struct
      console.log(index);
      if (index > 0) {
        p.name = `${lowerCaseToCamelCase(r.name.split(":")[1])}${p.name}`;
      }

      const data = InterfaceToGoStruct(p);

      //   Check if data contains *types. since we need to import it
      if (data.includes("*types.")) {
        structs = [
          ...structs,
          "import (",
          '  "github.com/respectZ/glowstone/bp/types"',
          ")",
          "",
        ];
      }
      structs.push(data);
      index++;
    });

    structs.push("\n");
    const data = structs.join("\n");
    const fname = r.name.split(":")[1];
    const dst = `./bp/item/component/${fname}.go`;
    write(dst, data);
  }
})();
