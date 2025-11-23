import { useEffect, useState } from "react";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Plus, Edit } from "lucide-react";

import { useNavigate } from "react-router-dom";

interface Journey {
    uuid: string;
    projectUUID: string;
    entityJourneys: EntityJourney[];
}

interface EntityJourney {
    uuid: string;
    entityName: string;
    operations: Operation[];
}

interface Operation {
    uuid: string;
    type: string;
    name: string;
    description: string;
}

interface JourneyListProps {
    projectId: string;
}

export function JourneyList({ projectId }: JourneyListProps) {
    const navigate = useNavigate();
    const [journeys, setJourneys] = useState<Journey[]>([]);
    const [loading, setLoading] = useState(false);

    const fetchJourneys = async () => {
        setLoading(true);
        try {
            const res = await api.post<{ items: Journey[] }>("/journeys/get-by-filter", {
                pageNumber: 1,
                pageSize: 100,
                dynamicFilter: {
                    field: "project_uuid",
                    operator: "eq",
                    value: projectId,
                },
            });
            setJourneys(res.items || []);
        } catch (error) {
            console.error("Failed to fetch journeys", error);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        if (projectId) {
            fetchJourneys();
        }
    }, [projectId]);

    if (loading) {
        return <div>Loading journeys...</div>;
    }

    if (journeys.length === 0) {
        return (
            <Card className="p-12 text-center">
                <div className="text-muted-foreground">
                    <p className="mb-4">No journeys configured yet</p>
                    <Button onClick={() => navigate(`/projects/${projectId}/journeys/builder`)}>
                        <Plus className="mr-2 h-4 w-4" /> Configure Journeys
                    </Button>
                </div>
            </Card>
        );
    }

    return (
        <div className="space-y-6">
            {journeys.map((journey) => (
                <div key={journey.uuid} className="space-y-4">
                    {journey.entityJourneys?.map((ej) => (
                        <div key={ej.uuid} className="space-y-3">
                            <h3 className="text-lg font-semibold">Entity: {ej.entityName}</h3>
                            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                                {ej.operations?.map((op) => (
                                    <Card key={op.uuid} className="hover:shadow-md transition-shadow cursor-pointer">
                                        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                                            <CardTitle className="text-base font-medium">
                                                {op.name}
                                            </CardTitle>
                                            <div className="flex gap-2">
                                                <Button variant="ghost" size="sm">
                                                    <Edit className="h-4 w-4" />
                                                </Button>
                                            </div>
                                        </CardHeader>
                                        <CardContent>
                                            <div className="flex items-center gap-2 mb-2">
                                                <span className="text-xs font-semibold bg-blue-100 text-blue-800 px-2 py-0.5 rounded">
                                                    {op.type}
                                                </span>
                                            </div>
                                            <p className="text-sm text-muted-foreground line-clamp-2">
                                                {op.description || "No description"}
                                            </p>
                                        </CardContent>
                                    </Card>
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            ))}
        </div>
    );
}
