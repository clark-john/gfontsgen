import { access, readFile } from "fs/promises";

const 
  MISSING_REQUIRED = "Missing required property: "
  ,VARIANTS_MINIMUM = "Variants must have at least one item"
;

export class ConfigProvider {
  private data?: Partial<Config>;
  private isValid = false;

  async useFile(file: string): Promise<Config> {
    try {
      await access(file);
    } catch (e) {
      this.throwError("Config file not found");
    }

    const data = await readFile(file, "utf-8");

    if (!this.isValidJSON(data))
      this.throwError("File has invalid JSON");
    
    this.data = JSON.parse(data);

    this.validate();
    this.isValid = true;

    return this.data as Config;
  }

  hasConfig() {
    return this.isValid;
  }

  private validate() {
    const opItems = this.data?.options;

    if (!opItems)
      return this.throwError(MISSING_REQUIRED + "options");

    if (!opItems.length) 
      this.throwError("Options must have at least one item.");

    /* error map impl from golang version */
    const errorMap: ErrorMap = {};

    opItems.forEach((item, index) => {
      const ff = item.fontFamily;
      const varnts = item.variants;

      if (!varnts && !ff)
        errorMap[index] = ["fontFamily", "variants"].map(x => MISSING_REQUIRED + x);
      else if (!ff)
        errorMap[index] = [MISSING_REQUIRED + "fontFamily"];
      else if (!varnts)
        errorMap[index] = [MISSING_REQUIRED + "variants"];

      if (varnts)
        if (!varnts.length)
          errorMap[index].push(VARIANTS_MINIMUM);
    });

    if (Object.keys(errorMap).length)
      this.printErrorMap(errorMap);
  }

  private printErrorMap(errorMap: ErrorMap) {
    const lines = ["Following option items are invalid:"];
    
    for (const [index, errs] of Object.entries(errorMap)) {
      lines.push(`  At index ${index}:`);
      for (const err of errs)
        lines.push("    " + err);
    }

    this.throwError(lines.join("\n"));
  }

  private throwError(message: string) {
    throw new Error(message);
  }

  private isValidJSON(text: string) {
    try {
      JSON.parse(text);
      return true;
    } catch (e) {
      return false;
    }
  }
}
