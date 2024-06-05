import { Command } from "commander";
import { CommandBase } from "./commands/index.js";
import { 
	GfontswebCommand, 
	ListCommand, 
	GenCommand 
} from "./commands/index.js";

const KEY_NAME = "GFONTSGEN_API_KEY";

if (!process.env.GFONTSGEN_API_KEY) {
	console.log("Set your environment variable", KEY_NAME, "to your API key you obtained from the Google Fonts API");
	console.log("Get your API key here: https://developers.google.com/fonts/docs/developer_api#APIKey");
	process.exit(-1);
}

const cmd = new Command();

const commands: CommandBase[] = [
	new ListCommand(),
	new GfontswebCommand(),
	new GenCommand()
];

cmd
	.version("1.1.0")
	.showHelpAfterError();

for (const comm of commands)
	cmd.addCommand(
		comm.action(comm.handle)
			.showHelpAfterError()
	);

cmd.parseAsync()
	.then()
	.catch(err => console.error(err));
