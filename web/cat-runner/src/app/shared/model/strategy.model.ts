export class StrategyBase {
    Name!: string;
    CreatedAt!: string;
    ModifiedAt!: string;
}

export class Strategy extends StrategyBase {
    HttpRequestRunners!: HttpRunner[]
    EnvironmentVariables!: EnvironmentVariable[]
}

export class HttpRunner {
    RequestName!: string;
    Http!: Http;
    VirtualUser!: VirtualUser;
}
  
export class Http {
    Protocol!: string;
    Host!: string;
    Port!: number;
    Path!: string;
    HttpMethod!: string;
    Timeout!: number;
    Headers!: KVParam[];
    Parameters!: KVParam[];
    Body?: Body;
}
  
export class KVParam {
    Key!: string;
    Value!: string;
}
  
export class Body {
    BodyFormat!: string;
    ContentText!: string;
}
  
export class VirtualUser {
    UsersAmount!: number;
    InteractionsAmount!: number;
    InteractionDelay!: number;
}

export class EnvironmentVariable {
    Name!:     string
	Value!:    string
	DataType!: string
}
