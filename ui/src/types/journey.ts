export interface Journey {
    uuid: string;
    projectUuid: string;
    programmingLanguage: string;
    blueprintId: string;
    entityJourneys: EntityJourney[];
}

export interface EntityJourney {
    uuid: string;
    entityId: string;
    entityName: string;
    operations: Operation[];
}

export interface Operation {
    uuid: string;
    type: string;
    name: string;
    description: string;
    frontendJourney: any[];
    backendJourney: JourneyStep[];
    filters?: Filter[];
    sort?: Sort[];
}

export interface JourneyStep {
    uuid: string;
    index: number;
    type: string;
    description?: string;
    fieldsInvolved?: FieldInvolved[];
    condition?: string;
    abortOnFail?: boolean;
    error?: string;
    curl?: string;
    sampleResponse?: string;
    retry?: boolean;
    retryCount?: number;
    retryInterval?: number;
    retryConditions?: RetryCondition[];
    responseActions?: ResponseAction[];
    dbAction?: string;
    channels?: string[];
    message?: string;
    recipients?: string[];
}

export interface FieldInvolved {
    uuid: string;
    id: string;
    name: string;
    source?: string;
}

export interface RetryCondition {
    uuid: string;
    condition: string;
    error: string;
}

export interface ResponseAction {
    uuid: string;
    index: number;
    type: string;
    fieldId?: string;
    value?: string;
    description?: string;
    fieldsInvolved?: ResFieldInvolved[];
    condition?: string;
    abortOnFail?: boolean;
    error?: string;
    nestedResponseAction?: ResponseAction;
}

export interface ResFieldInvolved {
    uuid: string;
    id: string;
    name: string;
    source?: string;
}

export interface Filter {
    uuid: string;
    name: string;
    type: string;
    fieldId: string;
    maxRange?: Range;
    minRange?: Range;
    error?: string;
    operator?: string;
}

export interface Range {
    uuid: string;
    value: number;
    unit: string;
}

export interface Sort {
    uuid: string;
    fieldId: string;
}
