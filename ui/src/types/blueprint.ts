export interface Blueprint {
    uuid: string;
    standardName: string;
    type: string;
    description: string;
    functionalities?: Functionality[];
}

export interface Functionality {
    uuid: string;
    category: string;
    type: string;
    provider: string;
    implementsGenerics: boolean;
    filePathsCSV: string;
    operations?: FunctionalOperation[];
}

export interface FunctionalOperation {
    uuid: string;
    name: string;
    description: string;
}
