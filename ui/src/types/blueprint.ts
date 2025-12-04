export interface Blueprint {
    uuid: string;
    standardName: string;
    type: string;
    description: string;
    functionalities?: Functionality[];
    libraries?: Library[];
}

export interface Library {
    uuid: string;
    standardName: string;
    version: string;
    description: string;
    repositoryURL: string;
    namespace: string;
    exposedFunctionalities?: LibraryFunctionality[];
}

export interface LibraryFunctionality {
    uuid: string;
    name: string;
    type: string;
    description: string;
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
