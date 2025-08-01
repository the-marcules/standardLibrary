export namespace cryptokit {
	
	export class JWK {
	    kty?: string;
	    crv?: string;
	    x?: string;
	    y?: string;
	    kid?: string;
	    alg?: string;
	    use?: string;
	
	    static createFrom(source: any = {}) {
	        return new JWK(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.kty = source["kty"];
	        this.crv = source["crv"];
	        this.x = source["x"];
	        this.y = source["y"];
	        this.kid = source["kid"];
	        this.alg = source["alg"];
	        this.use = source["use"];
	    }
	}
	export class Key {
	    kid?: string;
	
	    static createFrom(source: any = {}) {
	        return new Key(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.kid = source["kid"];
	    }
	}

}

