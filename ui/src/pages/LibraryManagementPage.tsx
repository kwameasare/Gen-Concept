import { useEffect, useState } from "react";
import { Plus, Trash2, ExternalLink } from "lucide-react";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import type { Library, LibraryFunctionality } from "@/types/blueprint";
import { LIBRARY_FUNCTIONALITY_TYPES } from "@/constants/library";

export default function LibraryManagementPage() {
    const [libraries, setLibraries] = useState<Library[]>([]);
    const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    const [newLibrary, setNewLibrary] = useState<Partial<Library>>({
        standardName: "",
        version: "",
        description: "",
        repositoryURL: "",
        namespace: "",
        exposedFunctionalities: [],
    });

    const [newFunctionality, setNewFunctionality] = useState<Partial<LibraryFunctionality>>({
        name: "",
        type: "",
        description: "",
    });

    useEffect(() => {
        fetchLibraries();
    }, []);

    const fetchLibraries = async () => {
        try {
            const response: any = await api.post("/libraries/get-by-filter", {
                pageNumber: 1,
                pageSize: 100,
            });
            setLibraries(response.items || []);
        } catch (error) {
            console.error("Error fetching libraries:", error);
        }
    };

    const handleCreateLibrary = async () => {
        if (!newLibrary.standardName || !newLibrary.version) {
            return;
        }

        setIsLoading(true);
        try {
            await api.post("/libraries", newLibrary);
            setIsCreateDialogOpen(false);
            setNewLibrary({
                standardName: "",
                version: "",
                description: "",
                repositoryURL: "",
                namespace: "",
                exposedFunctionalities: [],
            });
            fetchLibraries();
        } catch (error) {
            console.error("Error creating library:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const handleDeleteLibrary = async (uuid: string) => {
        try {
            await api.delete(`/libraries/${uuid}`);
            fetchLibraries();
        } catch (error) {
            console.error("Error deleting library:", error);
        }
    };

    const addFunctionality = () => {
        if (!newFunctionality.name || !newFunctionality.type) return;

        setNewLibrary({
            ...newLibrary,
            exposedFunctionalities: [
                ...(newLibrary.exposedFunctionalities || []),
                { ...newFunctionality, uuid: crypto.randomUUID() } as LibraryFunctionality,
            ],
        });

        setNewFunctionality({ name: "", type: "", description: "" });
    };

    const removeFunctionality = (index: number) => {
        setNewLibrary({
            ...newLibrary,
            exposedFunctionalities: newLibrary.exposedFunctionalities?.filter((_, i) => i !== index),
        });
    };

    return (
        <div className="min-h-screen bg-background p-8">
            <div className="max-w-6xl mx-auto space-y-8">
                <div className="flex items-center justify-between">
                    <div>
                        <h1 className="text-3xl font-bold tracking-tight">Library Management</h1>
                        <p className="text-muted-foreground">
                            Manage your organization's internal libraries
                        </p>
                    </div>

                    <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
                        <DialogTrigger asChild>
                            <Button>
                                <Plus className="mr-2 h-4 w-4" />
                                Create Library
                            </Button>
                        </DialogTrigger>
                        <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
                            <DialogHeader>
                                <DialogTitle>Create New Library</DialogTitle>
                                <DialogDescription>
                                    Add a new internal library to your catalog
                                </DialogDescription>
                            </DialogHeader>

                            <div className="space-y-4">
                                <div className="space-y-2">
                                    <Label>Library Name *</Label>
                                    <Input
                                        placeholder="e.g., MyOrg.Common.Utilities"
                                        value={newLibrary.standardName}
                                        onChange={(e) =>
                                            setNewLibrary({ ...newLibrary, standardName: e.target.value })
                                        }
                                    />
                                </div>

                                <div className="grid grid-cols-2 gap-4">
                                    <div className="space-y-2">
                                        <Label>Version *</Label>
                                        <Input
                                            placeholder="e.g., 2.1.0"
                                            value={newLibrary.version}
                                            onChange={(e) =>
                                                setNewLibrary({ ...newLibrary, version: e.target.value })
                                            }
                                        />
                                    </div>

                                    <div className="space-y-2">
                                        <Label>Namespace</Label>
                                        <Input
                                            placeholder="e.g., MyOrg.Common"
                                            value={newLibrary.namespace}
                                            onChange={(e) =>
                                                setNewLibrary({ ...newLibrary, namespace: e.target.value })
                                            }
                                        />
                                    </div>
                                </div>

                                <div className="space-y-2">
                                    <Label>Repository URL</Label>
                                    <Input
                                        placeholder="https://github.com/myorg/library"
                                        value={newLibrary.repositoryURL}
                                        onChange={(e) =>
                                            setNewLibrary({ ...newLibrary, repositoryURL: e.target.value })
                                        }
                                    />
                                </div>

                                <div className="space-y-2">
                                    <Label>Description</Label>
                                    <Textarea
                                        placeholder="What does this library do?"
                                        value={newLibrary.description}
                                        onChange={(e) =>
                                            setNewLibrary({ ...newLibrary, description: e.target.value })
                                        }
                                    />
                                </div>

                                <div className="border-t pt-4">
                                    <h3 className="font-semibold mb-4">Exposed Functionalities</h3>

                                    {/* Add Functionality Form */}
                                    <Card className="mb-4">
                                        <CardContent className="pt-4">
                                            <div className="grid grid-cols-2 gap-4 mb-4">
                                                <div className="space-y-2">
                                                    <Label>Functionality Name</Label>
                                                    <Input
                                                        placeholder="e.g., DateHelpers"
                                                        value={newFunctionality.name}
                                                        onChange={(e) =>
                                                            setNewFunctionality({
                                                                ...newFunctionality,
                                                                name: e.target.value,
                                                            })
                                                        }
                                                    />
                                                </div>

                                                <div className="space-y-2">
                                                    <Label>Type</Label>
                                                    <Select
                                                        value={newFunctionality.type}
                                                        onValueChange={(value) =>
                                                            setNewFunctionality({
                                                                ...newFunctionality,
                                                                type: value,
                                                            })
                                                        }
                                                    >
                                                        <SelectTrigger>
                                                            <SelectValue placeholder="Select type" />
                                                        </SelectTrigger>
                                                        <SelectContent>
                                                            {LIBRARY_FUNCTIONALITY_TYPES.map((type) => (
                                                                <SelectItem key={type} value={type}>
                                                                    {type}
                                                                </SelectItem>
                                                            ))}
                                                        </SelectContent>
                                                    </Select>
                                                </div>
                                            </div>

                                            <div className="space-y-2 mb-4">
                                                <Label>Description</Label>
                                                <Input
                                                    placeholder="What does it do?"
                                                    value={newFunctionality.description}
                                                    onChange={(e) =>
                                                        setNewFunctionality({
                                                            ...newFunctionality,
                                                            description: e.target.value,
                                                        })
                                                    }
                                                />
                                            </div>

                                            <Button
                                                type="button"
                                                variant="outline"
                                                size="sm"
                                                onClick={addFunctionality}
                                            >
                                                <Plus className="mr-2 h-4 w-4" />
                                                Add Functionality
                                            </Button>
                                        </CardContent>
                                    </Card>

                                    {/* List of Added Functionalities */}
                                    {newLibrary.exposedFunctionalities &&
                                        newLibrary.exposedFunctionalities.length > 0 && (
                                            <div className="space-y-2">
                                                {newLibrary.exposedFunctionalities.map((func, index) => (
                                                    <Card key={index}>
                                                        <CardContent className="pt-4">
                                                            <div className="flex items-start justify-between">
                                                                <div className="flex-1">
                                                                    <p className="font-medium">{func.name}</p>
                                                                    <p className="text-sm text-muted-foreground">
                                                                        {func.type} - {func.description}
                                                                    </p>
                                                                </div>
                                                                <Button
                                                                    variant="ghost"
                                                                    size="sm"
                                                                    onClick={() => removeFunctionality(index)}
                                                                >
                                                                    <Trash2 className="h-4 w-4" />
                                                                </Button>
                                                            </div>
                                                        </CardContent>
                                                    </Card>
                                                ))}
                                            </div>
                                        )}
                                </div>
                            </div>

                            <DialogFooter>
                                <Button
                                    variant="outline"
                                    onClick={() => setIsCreateDialogOpen(false)}
                                    disabled={isLoading}
                                >
                                    Cancel
                                </Button>
                                <Button onClick={handleCreateLibrary} disabled={isLoading}>
                                    Create Library
                                </Button>
                            </DialogFooter>
                        </DialogContent>
                    </Dialog>
                </div>

                {/* Libraries List */}
                <div className="grid gap-4">
                    {libraries.map((library) => (
                        <Card key={library.uuid}>
                            <CardHeader>
                                <div className="flex items-start justify-between">
                                    <div className="flex-1">
                                        <CardTitle className="text-xl">{library.standardName}</CardTitle>
                                        <p className="text-sm text-muted-foreground mt-1">
                                            Version {library.version} • {library.namespace}
                                        </p>
                                    </div>
                                    <Button
                                        variant="destructive"
                                        size="sm"
                                        onClick={() => handleDeleteLibrary(library.uuid)}
                                    >
                                        <Trash2 className="h-4 w-4" />
                                    </Button>
                                </div>
                            </CardHeader>
                            <CardContent>
                                <p className="text-sm mb-4">{library.description}</p>

                                {library.repositoryURL && (
                                    <a
                                        href={library.repositoryURL}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="text-sm text-blue-600 dark:text-blue-400 hover:underline inline-flex items-center mb-4"
                                    >
                                        <ExternalLink className="h-3 w-3 mr-1" />
                                        Repository
                                    </a>
                                )}

                                {library.exposedFunctionalities &&
                                    library.exposedFunctionalities.length > 0 && (
                                        <div className="mt-4">
                                            <h4 className="font-medium mb-2">Exposed Functionalities:</h4>
                                            <div className="grid grid-cols-2 gap-2">
                                                {library.exposedFunctionalities.map((func) => (
                                                    <div
                                                        key={func.uuid}
                                                        className="p-2 border rounded text-sm"
                                                    >
                                                        <p className="font-medium">{func.name}</p>
                                                        <p className="text-xs text-muted-foreground">
                                                            {func.type}
                                                        </p>
                                                    </div>
                                                ))}
                                            </div>
                                        </div>
                                    )}
                            </CardContent>
                        </Card>
                    ))}

                    {libraries.length === 0 && (
                        <Card>
                            <CardContent className="py-12 text-center text-muted-foreground">
                                No libraries yet. Create your first internal library to get started.
                            </CardContent>
                        </Card>
                    )}
                </div>
            </div>
        </div>
    );
}
