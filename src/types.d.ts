declare namespace NodeJS {
	interface ProcessEnv {
		GFONTSGEN_API_KEY?: string;
	}
}

type Classification = "display" | "handwriting" | "mono" | "monospace";
type DecorativeStroke = "sans-serif" | "serif" | "slab-serif";

interface FontItem {
	family: string;
	variants: string[];
	subsets: string[];
	version: string;
	lastModified: string;
	files: Record<string, string>;
	category: Classification & DecorativeStroke;
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
