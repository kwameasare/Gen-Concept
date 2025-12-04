export const PACKAGE_MANAGERS = [
    "NPM",
    "NUGET",
    "PIP",
    "MAVEN",
    "GO_MODULES",
    "COMPOSER",
    "GEM",
    "CARGO",
    "GRADLE",
] as const;

export type PackageManager = typeof PACKAGE_MANAGERS[number];

export const LIBRARY_FUNCTIONALITY_TYPES = [
    "UTILITY",
    "SERVICE",
    "HELPER",
    "EXTENSION",
    "MIDDLEWARE",
    "VALIDATOR",
    "MAPPER",
    "HANDLER",
] as const;

export type LibraryFunctionalityType = typeof LIBRARY_FUNCTIONALITY_TYPES[number];
