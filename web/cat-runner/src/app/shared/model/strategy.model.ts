export interface StrategyBase {
    Name: string;
    CreatedAt: string;
    ModifiedAt: string;
}

export interface Strategy extends StrategyBase {
    HttpRequestRunners: HttpRunner[]
    EnvironmentVariables: EnvironmentVariable[]
}

export interface HttpRunner {
    RequestName: string;
    Http: Http;
    VirtualUser: VirtualUser;
}

export interface Http {
    Protocol: string;
    Host: string;
    Port: number;
    Path: string;
    HttpMethod: string;
    Timeout: number;
    Headers: KVParam[];
    Parameters: KVParam[];
    Body: Body;
}
  
export interface KVParam {
    Key: string;
    Value: string;
}
  
export interface Body {
    BodyFormat: string;
    ContentText: string;
}

export interface VirtualUser {
    UsersAmount: number;
    InteractionsAmount: number;
    InteractionDelay: number;
}

export interface EnvironmentVariable {
    Name:     string
	Value:    string
	DataType: string
}
