export namespace agentcli {
	
	export class CLI {
	    id: string;
	    label: string;
	    binary: string;
	    available: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CLI(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	        this.binary = source["binary"];
	        this.available = source["available"];
	    }
	}
	export class Request {
	    cliId: string;
	    action: string;
	    testName: string;
	    status: string;
	    failure: string;
	    story: string;
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cliId = source["cliId"];
	        this.action = source["action"];
	        this.testName = source["testName"];
	        this.status = source["status"];
	        this.failure = source["failure"];
	        this.story = source["story"];
	    }
	}

}

export namespace main {
	
	export class ActionResult {
	    started: boolean;
	    blocked: boolean;
	    failed: boolean;
	    message: string;
	    usedPath: string;
	    output: string;
	    exitCode: number;
	
	    static createFrom(source: any = {}) {
	        return new ActionResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.started = source["started"];
	        this.blocked = source["blocked"];
	        this.failed = source["failed"];
	        this.message = source["message"];
	        this.usedPath = source["usedPath"];
	        this.output = source["output"];
	        this.exitCode = source["exitCode"];
	    }
	}
	export class AgentState {
	    running: boolean;
	    finished: boolean;
	    stopped: boolean;
	    failed: boolean;
	    unavailable: boolean;
	    message: string;
	    output: string;
	    exitCode: number;
	    cliId: string;
	    cliLabel: string;
	    cliBinary: string;
	    action: string;
	    testName: string;
	    pendingDecision: boolean;
	    proposedChanges: projectsnap.Change[];
	    accepted: boolean;
	    discarded: boolean;
	    canRerun: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AgentState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.running = source["running"];
	        this.finished = source["finished"];
	        this.stopped = source["stopped"];
	        this.failed = source["failed"];
	        this.unavailable = source["unavailable"];
	        this.message = source["message"];
	        this.output = source["output"];
	        this.exitCode = source["exitCode"];
	        this.cliId = source["cliId"];
	        this.cliLabel = source["cliLabel"];
	        this.cliBinary = source["cliBinary"];
	        this.action = source["action"];
	        this.testName = source["testName"];
	        this.pendingDecision = source["pendingDecision"];
	        this.proposedChanges = this.convertValues(source["proposedChanges"], projectsnap.Change);
	        this.accepted = source["accepted"];
	        this.discarded = source["discarded"];
	        this.canRerun = source["canRerun"];
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
	export class EnvStatus {
	    loaded: boolean;
	    error: string;
	    path: string;
	    vars: Record<string, string>;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new EnvStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.loaded = source["loaded"];
	        this.error = source["error"];
	        this.path = source["path"];
	        this.vars = source["vars"];
	        this.count = source["count"];
	    }
	}
	export class RunSettings {
	    projectPath: string;
	    command: string;
	
	    static createFrom(source: any = {}) {
	        return new RunSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.projectPath = source["projectPath"];
	        this.command = source["command"];
	    }
	}
	export class SuiteState {
	    running: boolean;
	    finished: boolean;
	    stopped: boolean;
	    failed: boolean;
	    message: string;
	    output: string;
	    exitCode: number;
	    tests: suiteparse.Test[];
	    counts: suiteparse.Counts;
	
	    static createFrom(source: any = {}) {
	        return new SuiteState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.running = source["running"];
	        this.finished = source["finished"];
	        this.stopped = source["stopped"];
	        this.failed = source["failed"];
	        this.message = source["message"];
	        this.output = source["output"];
	        this.exitCode = source["exitCode"];
	        this.tests = this.convertValues(source["tests"], suiteparse.Test);
	        this.counts = this.convertValues(source["counts"], suiteparse.Counts);
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

export namespace projectsnap {
	
	export class Change {
	    path: string;
	    kind: string;
	
	    static createFrom(source: any = {}) {
	        return new Change(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.kind = source["kind"];
	    }
	}

}

export namespace suiteparse {
	
	export class Counts {
	    finished: number;
	    running: number;
	    waiting: number;
	
	    static createFrom(source: any = {}) {
	        return new Counts(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.finished = source["finished"];
	        this.running = source["running"];
	        this.waiting = source["waiting"];
	    }
	}
	export class Test {
	    name: string;
	    status: string;
	    duration?: string;
	    story?: string;
	    failure?: string;
	    hasPrevious?: boolean;
	    previousStatus?: string;
	    previousDuration?: string;
	    previousFailure?: string;
	
	    static createFrom(source: any = {}) {
	        return new Test(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.status = source["status"];
	        this.duration = source["duration"];
	        this.story = source["story"];
	        this.failure = source["failure"];
	        this.hasPrevious = source["hasPrevious"];
	        this.previousStatus = source["previousStatus"];
	        this.previousDuration = source["previousDuration"];
	        this.previousFailure = source["previousFailure"];
	    }
	}

}

