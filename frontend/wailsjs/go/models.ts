export namespace controller {
	
	export class ImagesRouteBase64 {
	    base64?: string;
	    url: string;
	    url_server: string;
	    url_server_path: string;
	
	    static createFrom(source: any = {}) {
	        return new ImagesRouteBase64(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.base64 = source["base64"];
	        this.url = source["url"];
	        this.url_server = source["url_server"];
	        this.url_server_path = source["url_server_path"];
	    }
	}

}

export namespace models {
	
	export class Folder {
	    name: string;
	    description: string;
	    rute: string;
	    termination_image: string;
	
	    static createFrom(source: any = {}) {
	        return new Folder(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.rute = source["rute"];
	        this.termination_image = source["termination_image"];
	    }
	}
	export class JsonResponse {
	    value: any;
	    text: string;
	
	    static createFrom(source: any = {}) {
	        return new JsonResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.text = source["text"];
	    }
	}

}

