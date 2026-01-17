export interface Team {
    uuid: string;
    name: string;
    description: string;
    organizationID: number;
}

export interface CreateTeamRequest {
    name: string;
    description: string;
    organizationID: number;
}

export interface UpdateTeamRequest {
    name?: string;
    description?: string;
}
