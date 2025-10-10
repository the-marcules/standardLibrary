export namespace fileOperation {
	
	export class CamMod {
	    Name: string;
	    Software: string;
	
	    static createFrom(source: any = {}) {
	        return new CamMod(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Software = source["Software"];
	    }
	}
	export class ImageResolution {
	    X: number;
	    Y: number;
	
	    static createFrom(source: any = {}) {
	        return new ImageResolution(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.X = source["X"];
	        this.Y = source["Y"];
	    }
	}
	export class LocLatLong {
	    Latitude: number;
	    Longitude: number;
	
	    static createFrom(source: any = {}) {
	        return new LocLatLong(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Latitude = source["Latitude"];
	        this.Longitude = source["Longitude"];
	    }
	}
	export class ExifInfo {
	    // Go type: LocLatLong
	    Location: any;
	    // Go type: time
	    DateTaken: any;
	    // Go type: ImageResolution
	    Size: any;
	    // Go type: CamMod
	    CameraModel: any;
	    ExifIFDPointer: string;
	    DateTimeDigitized: string;
	    Orientation: string;
	    YResolution: string;
	    PixelYDimension: string;
	    Software: string;
	    DateTime: string;
	    ColorSpace: string;
	    XResolution: string;
	    PixelXDimension: string;
	
	    static createFrom(source: any = {}) {
	        return new ExifInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Location = this.convertValues(source["Location"], null);
	        this.DateTaken = this.convertValues(source["DateTaken"], null);
	        this.Size = this.convertValues(source["Size"], null);
	        this.CameraModel = this.convertValues(source["CameraModel"], null);
	        this.ExifIFDPointer = source["ExifIFDPointer"];
	        this.DateTimeDigitized = source["DateTimeDigitized"];
	        this.Orientation = source["Orientation"];
	        this.YResolution = source["YResolution"];
	        this.PixelYDimension = source["PixelYDimension"];
	        this.Software = source["Software"];
	        this.DateTime = source["DateTime"];
	        this.ColorSpace = source["ColorSpace"];
	        this.XResolution = source["XResolution"];
	        this.PixelXDimension = source["PixelXDimension"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ExifResponse {
	    ExifInfo: ExifInfo;
	    FilePath: string;
	
	    static createFrom(source: any = {}) {
	        return new ExifResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ExifInfo = this.convertValues(source["ExifInfo"], ExifInfo);
	        this.FilePath = source["FilePath"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

