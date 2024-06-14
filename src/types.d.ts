declare namespace NodeJS {
	interface ProcessEnv {
		GFONTSGEN_API_KEY?: string;
	}
}

type Variant = 
	"200" | "200i" |
	"300" | "300i" |
	"500" | "500i" |
	"600" | "600i" |
	"700" | "700i" |
	"800" | "800i" |
	"900" | "900i" |
	"regular" | "italic"
;

interface OptionItem {
	fontFamily: string;
	variants: Variant[];
}

interface Config {
	copy?: boolean;
	woff?: boolean;
	toCssImport?: boolean;
	deleteFontDirBeforeDownload?: boolean;
	outputPath?: string;
	options: OptionItem[]
}

type Classification = "display" | "handwriting" | "mono" | "monospace";
type DecorativeStroke = "sans-serif" | "serif" | "slab-serif";

type ErrorMap = Record<number, string[]>;

interface FontItem {
	family: string;
	variants: string[];
	subsets: string[];
	version: string;
	lastModified: string;
	files: Record<string, string>;
	category: Classification | DecorativeStroke;
	kind: string;
	menu: string;
}

interface FullResponse {
	kind: string;
	items: FontItem[];
}

interface ErrorResponse {
	error: ErrorObject;
}

interface ErrorObject {
	code: number;
	message: string;
	errors: string[];
	status: string;
	details?: ErrorDetail[];
}

interface ErrorDetail {
	type: string;
	reason: string;
	domain: string;
	metadata: {
		service: string;
	};
}

interface ErrorItem {
	message: string;
	detail: string;
	reason: string;
}
