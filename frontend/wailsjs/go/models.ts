export namespace app {
	
	export class Episode {
	    aid: number;
	    bvid: string;
	    cid: number;
	    title: string;
	    duration: string;
	    index: number;
	    selected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Episode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.aid = source["aid"];
	        this.bvid = source["bvid"];
	        this.cid = source["cid"];
	        this.title = source["title"];
	        this.duration = source["duration"];
	        this.index = source["index"];
	        this.selected = source["selected"];
	    }
	}
	export class Stat {
	    view: string;
	    danmaku: string;
	    like: string;
	    coin: string;
	    favorite: string;
	    share: string;
	    reply: string;
	
	    static createFrom(source: any = {}) {
	        return new Stat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.view = source["view"];
	        this.danmaku = source["danmaku"];
	        this.like = source["like"];
	        this.coin = source["coin"];
	        this.favorite = source["favorite"];
	        this.share = source["share"];
	        this.reply = source["reply"];
	    }
	}
	export class VideoInfo {
	    bvid: string;
	    title: string;
	    desc: string;
	    pic: string;
	    pubdate: string;
	    owner_name: string;
	    owner_face: string;
	    stat: Stat;
	    eps: Episode[];
	
	    static createFrom(source: any = {}) {
	        return new VideoInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bvid = source["bvid"];
	        this.title = source["title"];
	        this.desc = source["desc"];
	        this.pic = source["pic"];
	        this.pubdate = source["pubdate"];
	        this.owner_name = source["owner_name"];
	        this.owner_face = source["owner_face"];
	        this.stat = this.convertValues(source["stat"], Stat);
	        this.eps = this.convertValues(source["eps"], Episode);
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

}

