import { Option } from "commander";
import { CommandBase } from "./base.js";
import { FontFilesCommand } from "./gen/fontfiles.js";
import { UrlCommand } from "./gen/url.js";

export const configOption = 
  new Option("--config <file>", "Url/font gen config file to use");

export const variantsOption = 
  new Option(
    "-v, --variants <variants>", 
    "Generate with specific variants (e.g. 400, 600, 200i, etc)"
  );

export interface CommonGenOptions {
  variants: string;
  config: string;
}

export class GenCommand extends CommandBase {
  constructor() {
    super("gen");
    
    for (const command of [
      new FontFilesCommand(),
      new UrlCommand()
    ])
      this.addCommand(
        command
          .action(command.handle)
          .showHelpAfterError()
      );

    this.description("Generate font URL/font files");
  }

  async handle() {
    if (!this.args.length)
      return this.help();

    if (!this.commands.map(x => x.name()).includes(this.args[0]))
      this.error(`error: Unknown command "${this.args[0]}"`);
  }
}
