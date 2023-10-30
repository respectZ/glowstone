import {
  DOMParser,
  Element,
} from "https://deno.land/x/deno_dom@v0.1.38/deno-dom-wasm.ts";

interface BaseDescription {
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
  decimal: "float64",
  boolean: "bool",
  array: "[]interface{}",
  list: "[]interface{}",
  "vector [x,y,z]": "[]float64",
  "vector [a, b, c]": "[]float64",
  event: "string",
  "integer range": "[]int",
  trigger: "*types.Trigger",
  "json object": "interface{}",
  enumerator: "string", // minecraft:wearable
  "": "int", // minecraft:max_stack_size
  "--": "interface{}", // minecraft:enchantable
};

const URL =
  "https://learn.microsoft.com/en-us/minecraft/creator/reference/content/blockreference/examples/blockcomponents/blockcomponentslist";

function BaseDescription(tr: Element): BaseDescription {
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

function lowerCaseToCamelCase(str: string): string {
  let s = str.replace(/_([a-z])/g, function (g) {
    return g[1].toUpperCase();
  });
  s = s.charAt(0).toUpperCase() + s.slice(1);
  return s;
}

function write(dest: string, data: string): void {
  Deno.writeTextFileSync(dest, data);
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
        name = el.textContent?.trim().replaceAll(" Parameters", "") ?? "";
        break;
    }

    // Get fields
    try {
      const fields = TableToFields(table);
      const struct: GoStruct = {
        name,
        fields,
        description,
      };
      r.push(struct);
    } catch (e) {
      console.error(e);
    }
  });
  return r;
}

function TableToFields(table: Element): GoField[] {
  // Loop tbody and tr
  const _tbody = table.querySelector("tbody");
  if (!_tbody) throw new Error("Failed to find tbody");

  const thead = table.querySelector("thead");
  if (!thead) throw new Error("Failed to find thead");

  const thead_tr = thead.querySelector("tr");
  if (!thead_tr) throw new Error("Failed to find thead trs");

  const thead_tr_ths = thead_tr.querySelectorAll("th");
  if (!thead_tr_ths) throw new Error("Failed to find thead tr ths");

  // Loop through tr
  let nameIndex = -1;
  let typeIndex = -1;
  for (let i = 0; i < thead_tr_ths.length; i++) {
    const el = thead_tr_ths[i] as Element;
    // Check if innerText is Name
    if (el.textContent.trim().includes("Name")) {
      nameIndex = i + 1;
    } else if (el.textContent.trim().includes("Type")) {
      typeIndex = i + 1;
    }
  }
  if (nameIndex === -1) throw new Error("Failed to find nameIndex");
  if (typeIndex === -1) throw new Error("Failed to find typeIndex");

  // Loop through tr
  const _trs = _tbody.querySelectorAll("tr");
  const r: GoField[] = [];
  _trs.forEach((_tr) => {
    const _t = <Element>_tr;
    const name =
      _t.querySelector(`td:nth-child(${nameIndex})`)?.textContent?.trim() ?? "";
    const _type =
      _t.querySelector(`td:nth-child(${typeIndex})`)?.textContent?.trim() ?? "";

    // Check type from datamap
    const gtype = GoDataMap[_type.toLowerCase()] ?? _type;
    if (GoDataMap[_type.toLowerCase()] === undefined)
      console.error(`Unknown type: ${_type} in ${name}`);

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

(async () => {
  // Fetch
  const res = await fetch(URL);
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
    const tr = <Element>_tr;
    // First child is component name
    const componentName = tr.children[0].textContent;
    if (!componentName) throw new Error("Failed to find component name");
    // Fourth child is component description
    const description = tr.children[3].textContent;
    if (!description) throw new Error("Failed to find component description");

    const r = BaseDescription(tr);

    const ps = await GoStructFromURL(
      `https://learn.microsoft.com/en-us/minecraft/creator/reference/content/blockreference/examples/blockcomponents/${r.url}`
    );

    let structs: string[] = [
      "package component\n",
      "// Code generated by scrapper/block.ts;",
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
    const fname = componentName.split(":")[1];
    const dst = `./bp/block/component/${fname}.go`;
    write(dst, data);
  }
})();
