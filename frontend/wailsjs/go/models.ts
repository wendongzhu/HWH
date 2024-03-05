export namespace param {
	
	export class ExportParam {
	    model_group: string;
	    model_name: string;
	    program_number: string;
	    weld_time: string;
	    sampling_time: number;
	    extract_current: number;
	    extract_voltage: number;
	    extract_resistance: number;
	    peak_current: number;
	    peak_voltage: number;
	    peak_resistance: number;
	    archive_series: string;
	
	    static createFrom(source: any = {}) {
	        return new ExportParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.model_group = source["model_group"];
	        this.model_name = source["model_name"];
	        this.program_number = source["program_number"];
	        this.weld_time = source["weld_time"];
	        this.sampling_time = source["sampling_time"];
	        this.extract_current = source["extract_current"];
	        this.extract_voltage = source["extract_voltage"];
	        this.extract_resistance = source["extract_resistance"];
	        this.peak_current = source["peak_current"];
	        this.peak_voltage = source["peak_voltage"];
	        this.peak_resistance = source["peak_resistance"];
	        this.archive_series = source["archive_series"];
	    }
	}
	export class ExportListParam {
	    data: ExportParam[];
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new ExportListParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], ExportParam);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	
	export class FileNameParam {
	    value: string;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new FileNameParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.label = source["label"];
	    }
	}
	export class FileNameListParam {
	    data: FileNameParam[];
	
	    static createFrom(source: any = {}) {
	        return new FileNameListParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], FileNameParam);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	
	export class HWHArchiveParam {
	    spot_time: string;
	    model_group: string;
	    model_name: string;
	    type_id: string;
	    spot_number: string;
	    program_number: string;
	    spot_name: string;
	    control_mode: string;
	    set_current: string;
	    real_mean_current: string;
	    real_mean_voltage: string;
	    pre_time: string;
	    keep_time: string;
	    set_time: string;
	    real_time: string;
	    real_energy: string;
	    milling_cycle_counter: string;
	    milling_weld_spot_counter: string;
	    weld_spot_counter: string;
	    r_max: string;
	    q_value: string;
	    q_threshold: string;
	    spot_quality: string;
	    spatter: string;
	    gun_number: string;
	    gun_name: string;
	    error_code: string;
	    archive_series: string;
	    skt: number[];
	    ikv: number[];
	    ukv: number[];
	    rkv: number[];
	    qkv: number[];
	
	    static createFrom(source: any = {}) {
	        return new HWHArchiveParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.spot_time = source["spot_time"];
	        this.model_group = source["model_group"];
	        this.model_name = source["model_name"];
	        this.type_id = source["type_id"];
	        this.spot_number = source["spot_number"];
	        this.program_number = source["program_number"];
	        this.spot_name = source["spot_name"];
	        this.control_mode = source["control_mode"];
	        this.set_current = source["set_current"];
	        this.real_mean_current = source["real_mean_current"];
	        this.real_mean_voltage = source["real_mean_voltage"];
	        this.pre_time = source["pre_time"];
	        this.keep_time = source["keep_time"];
	        this.set_time = source["set_time"];
	        this.real_time = source["real_time"];
	        this.real_energy = source["real_energy"];
	        this.milling_cycle_counter = source["milling_cycle_counter"];
	        this.milling_weld_spot_counter = source["milling_weld_spot_counter"];
	        this.weld_spot_counter = source["weld_spot_counter"];
	        this.r_max = source["r_max"];
	        this.q_value = source["q_value"];
	        this.q_threshold = source["q_threshold"];
	        this.spot_quality = source["spot_quality"];
	        this.spatter = source["spatter"];
	        this.gun_number = source["gun_number"];
	        this.gun_name = source["gun_name"];
	        this.error_code = source["error_code"];
	        this.archive_series = source["archive_series"];
	        this.skt = source["skt"];
	        this.ikv = source["ikv"];
	        this.ukv = source["ukv"];
	        this.rkv = source["rkv"];
	        this.qkv = source["qkv"];
	    }
	}
	export class HWHArchiveDataListParam {
	    data: HWHArchiveParam[];
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new HWHArchiveDataListParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], HWHArchiveParam);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	
	export class ModelFileParam {
	    ArchiveFilePath: string;
	    ArchiveFileName: string;
	    ExportFilePath: string;
	    ExportFileName: string;
	    SamplingTime: number;
	
	    static createFrom(source: any = {}) {
	        return new ModelFileParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ArchiveFilePath = source["ArchiveFilePath"];
	        this.ArchiveFileName = source["ArchiveFileName"];
	        this.ExportFilePath = source["ExportFilePath"];
	        this.ExportFileName = source["ExportFileName"];
	        this.SamplingTime = source["SamplingTime"];
	    }
	}
	export class SelectDataParam {
	    file_name: string;
	    page_size: number;
	    current_page: number;
	
	    static createFrom(source: any = {}) {
	        return new SelectDataParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.file_name = source["file_name"];
	        this.page_size = source["page_size"];
	        this.current_page = source["current_page"];
	    }
	}

}

