// Functionality categories
export const FUNCTIONALITY_CATEGORIES = [
    "STANDARD",
    "CUSTOM",
] as const;

export type FunctionalityCategory = typeof FUNCTIONALITY_CATEGORIES[number];

// Functionality types
export const FUNCTIONALITY_TYPES = [
    "CACHE",
    "MESSAGE BROKER",
    "DATABASE",
    "LOGGING",
    "AUTHENTICATION",
    "EMAIL SERVICE",
    "SMS SERVICE",
    "PUSH NOTIFICATION",
    "OBJECT STORAGE",
    "HTTP CLIENT",
    "JOB SCHEDULER",
    "SEARCH ENGINE",
    "INPUT VALIDATION",
    "OBSERVABILITY",
] as const;

export type FunctionalityType = typeof FUNCTIONALITY_TYPES[number];

// Provider mappings based on type
export const PROVIDERS_BY_TYPE: Record<string, string[]> = {
    "CACHE": ["REDIS", "MEMCACHED", "IN-MEMORY"],
    "MESSAGE BROKER": ["RABBITMQ", "KAFKA", "AWS SQS", "AZURE SERVICE BUS", "REDIS"],
    "DATABASE": ["MYSQL", "POSTGRESQL", "MONGODB", "MSSQL", "ORACLE", "CASSANDRA"],
    "LOGGING": ["DEFAULT"],
    "AUTHENTICATION": ["JWT", "ACTIVE DIRECTORY", "OAUTH2", "SAML", "FIREBASE AUTH", "AWS COGNITO"],
    "EMAIL SERVICE": ["SENDGRID", "AWS SES", "MAILGUN", "SMTP", "POSTMARK"],
    "SMS SERVICE": ["TWILIO", "AWS SNS", "NEXMO", "MESSAGEBIRD"],
    "PUSH NOTIFICATION": ["FIREBASE", "AWS SNS", "ONESIGNAL", "PUSHWOOSH"],
    "OBJECT STORAGE": ["AZURE BLOB STORAGE", "AWS S3", "GOOGLE CLOUD STORAGE", "LOCAL STORAGE"],
    "HTTP CLIENT": ["DEFAULT"],
    "JOB SCHEDULER": ["DEFAULT"],
    "SEARCH ENGINE": ["ELASTICSEARCH", "ALGOLIA", "AZURE SEARCH", "SOLR"],
    "INPUT VALIDATION": ["MSISDN VALIDATOR", "EMAIL VALIDATOR", "REGEX VALIDATOR", "CUSTOM"],
    "OBSERVABILITY": ["OPENTELEMETRY", "DATADOG", "NEW RELIC", "APP INSIGHTS"],
};

// Helper function to get providers for a type
export function getProvidersForType(type: string): string[] {
    return PROVIDERS_BY_TYPE[type] || [];
}
