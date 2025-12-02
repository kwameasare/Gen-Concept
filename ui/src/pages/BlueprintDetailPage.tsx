import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { ArrowLeft, Edit, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { api } from "@/lib/api";
import BlueprintForm from "@/components/BlueprintForm";
import type { Blueprint } from "@/types/blueprint";
import type { BlueprintFormData } from "@/components/BlueprintForm";

export default function BlueprintDetailPage() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const [blueprint, setBlueprint] = useState<Blueprint | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [isEditing, setIsEditing] = useState(false);
    const [isSaving, setIsSaving] = useState(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (id) {
            fetchBlueprint();
        }
    }, [id]);

    const fetchBlueprint = async () => {
        try {
            setIsLoading(true);
            const data = await api.get<Blueprint>(`/blueprints/${id}`);
            setBlueprint(data);
            setError(null);
        } catch (err) {
            console.error("Failed to fetch blueprint", err);
            setError("Failed to load blueprint");
        } finally {
            setIsLoading(false);
        }
    };

    const handleUpdate = async (data: BlueprintFormData) => {
        try {
            setIsSaving(true);
            const updated = await api.put<Blueprint>(`/blueprints/${id}`, data);
            setBlueprint(updated);
            setIsEditing(false);
            setError(null);
        } catch (err) {
            console.error("Failed to update blueprint", err);
            setError("Failed to update blueprint");
        } finally {
            setIsSaving(false);
        }
    };

    if (isLoading) {
        return (
            <div className="min-h-screen bg-gray-50 p-8">
                <div className="max-w-4xl mx-auto">
                    <p>Loading...</p>
                </div>
            </div>
        );
    }

    if (error || !blueprint) {
        return (
            <div className="min-h-screen bg-gray-50 p-8">
                <div className="max-w-4xl mx-auto">
                    <Button variant="ghost" onClick={() => navigate("/dashboard")}>
                        <ArrowLeft className="mr-2 h-4 w-4" />
                        Back to Dashboard
                    </Button>
                    <p className="text-red-500 mt-4">{error || "Blueprint not found"}</p>
                </div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-gray-50 p-8">
            <div className="max-w-4xl mx-auto space-y-6">
                {/* Header */}
                <div className="flex items-center justify-between">
                    <Button variant="ghost" onClick={() => navigate("/dashboard")}>
                        <ArrowLeft className="mr-2 h-4 w-4" />
                        Back to Dashboard
                    </Button>
                    {!isEditing ? (
                        <Button onClick={() => setIsEditing(true)}>
                            <Edit className="mr-2 h-4 w-4" />
                            Edit
                        </Button>
                    ) : (
                        <Button variant="outline" onClick={() => setIsEditing(false)}>
                            <X className="mr-2 h-4 w-4" />
                            Cancel
                        </Button>
                    )}
                </div>

                {/* Content */}
                {isEditing ? (
                    <Card>
                        <CardHeader>
                            <CardTitle>Edit Blueprint</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <BlueprintForm
                                initialData={blueprint}
                                onSubmit={handleUpdate}
                                onCancel={() => setIsEditing(false)}
                                isLoading={isSaving}
                            />
                        </CardContent>
                    </Card>
                ) : (
                    <>
                        {/* Overview Card */}
                        <Card>
                            <CardHeader>
                                <CardTitle>{blueprint.standardName}</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-4">
                                <div>
                                    <h3 className="text-sm font-medium text-gray-500">Type</h3>
                                    <p className="mt-1">{blueprint.type}</p>
                                </div>
                                {blueprint.description && (
                                    <div>
                                        <h3 className="text-sm font-medium text-gray-500">
                                            Description
                                        </h3>
                                        <p className="mt-1 whitespace-pre-wrap">
                                            {blueprint.description}
                                        </p>
                                    </div>
                                )}
                            </CardContent>
                        </Card>

                        {/* Functionalities */}
                        {blueprint.functionalities && blueprint.functionalities.length > 0 && (
                            <div className="space-y-4">
                                <h2 className="text-2xl font-semibold">Functionalities</h2>
                                {blueprint.functionalities.map((func, index) => (
                                    <Card key={func.uuid || index}>
                                        <CardHeader>
                                            <CardTitle className="text-lg">
                                                {func.category} - {func.type}
                                            </CardTitle>
                                        </CardHeader>
                                        <CardContent className="space-y-4">
                                            <div className="grid grid-cols-2 gap-4">
                                                {func.provider && (
                                                    <div>
                                                        <h4 className="text-sm font-medium text-gray-500">
                                                            Provider
                                                        </h4>
                                                        <p className="mt-1">{func.provider}</p>
                                                    </div>
                                                )}
                                                <div>
                                                    <h4 className="text-sm font-medium text-gray-500">
                                                        Implements Generics
                                                    </h4>
                                                    <p className="mt-1">
                                                        {func.implementsGenerics ? "Yes" : "No"}
                                                    </p>
                                                </div>
                                                {func.filePathsCSV && (
                                                    <div className="col-span-2">
                                                        <h4 className="text-sm font-medium text-gray-500">
                                                            File Paths
                                                        </h4>
                                                        <p className="mt-1 text-sm font-mono">
                                                            {func.filePathsCSV}
                                                        </p>
                                                    </div>
                                                )}
                                            </div>

                                            {/* Operations */}
                                            {func.operations && func.operations.length > 0 && (
                                                <div className="border-t pt-4">
                                                    <h4 className="font-medium mb-3">Operations</h4>
                                                    <div className="space-y-2">
                                                        {func.operations.map((op, opIndex) => (
                                                            <div
                                                                key={op.uuid || opIndex}
                                                                className="p-3 bg-gray-50 rounded-md"
                                                            >
                                                                <h5 className="font-medium">
                                                                    {op.name}
                                                                </h5>
                                                                {op.description && (
                                                                    <p className="text-sm text-gray-600 mt-1">
                                                                        {op.description}
                                                                    </p>
                                                                )}
                                                            </div>
                                                        ))}
                                                    </div>
                                                </div>
                                            )}
                                        </CardContent>
                                    </Card>
                                ))}
                            </div>
                        )}
                    </>
                )}
            </div>
        </div>
    );
}
