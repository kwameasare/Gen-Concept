import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Plus, ArrowLeft, Trash2, Edit } from "lucide-react";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
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
import { Switch } from "@/components/ui/switch";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { JourneyList } from "@/components/JourneyList";

const entitySchema = z.object({
    entityName: z.string().min(2, "Entity name is required"),
    entityDescription: z.string().optional(),
    implementsRBAC: z.boolean().default(false),
    isAuthenticationRequired: z.boolean().default(false),
    implementsAudit: z.boolean().default(false),
    implementsChangeManagement: z.boolean().default(false),
    isReadOnly: z.boolean().default(false),
    isIndependentEntity: z.boolean().default(true),
    version: z.string().default("1.0"),
    isBackendOnly: z.boolean().default(false),
    preferredDB: z.string().default("Postgres"),
    modeOfDBInteraction: z.string().default("ORM"),
});

type EntityFormData = z.infer<typeof entitySchema>;

interface Project {
    uuid: string;
    projectName: string;
    projectDescription: string;
    projectType: string;
    isMultiTenant: boolean;
    isMultiLingual: boolean;
    entities: Entity[];
}

interface Entity {
    uuid: string;
    entityName: string;
    entityDescription: string;
    implementsRBAC: boolean;
    isAuthenticationRequired: boolean;
    implementsAudit: boolean;
    implementsChangeManagement: boolean;
    isReadOnly: boolean;
    isIndependentEntity: boolean;
    version: string;
    isBackendOnly: boolean;
    preferredDB: string;
    modeOfDBInteraction: string;
}

const DB_OPTIONS = ["Postgres", "MySQL", "MongoDB", "SQLite"];
const DB_INTERACTION_MODES = ["ORM", "StoredProcedures", "RawSQL"];

export default function ProjectDetailPage() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const [project, setProject] = useState<Project | null>(null);
    const [loadingProject, setLoadingProject] = useState(true);
    const [projectError, setProjectError] = useState<string | null>(null);
    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const {
        register,
        handleSubmit,
        reset,
        control,
        formState: { errors },
    } = useForm<EntityFormData>({
        resolver: zodResolver(entitySchema),
        defaultValues: {
            entityName: "",
            entityDescription: "",
            implementsRBAC: false,
            isAuthenticationRequired: false,
            implementsAudit: false,
            implementsChangeManagement: false,
            isReadOnly: false,
            isIndependentEntity: true,
            version: "1.0",
            isBackendOnly: false,
            preferredDB: "Postgres",
            modeOfDBInteraction: "ORM",
        },
    });

    const fetchProject = async () => {
        setLoadingProject(true);
        setProjectError(null);
        try {
            const res = await api.get<Project>(`/projects/${id}`);
            setProject(res);
        } catch (error: any) {
            console.error("Failed to fetch project", error);
            const errorMsg = error.response?.status === 429
                ? "Too many requests. Please wait a moment and try again."
                : "Failed to load project. Please try again.";
            setProjectError(errorMsg);
        } finally {
            setLoadingProject(false);
        }
    };

    useEffect(() => {
        if (id) {
            fetchProject();
        }
    }, [id]);

    const onSubmit = async (data: EntityFormData) => {
        setLoading(true);
        try {
            // Create entity under this project
            await api.post(`/entities/`, {
                ...data,
                projectId: id,
            });
            setOpen(false);
            reset();
            fetchProject();
        } catch (error) {
            console.error("Failed to create entity", error);
        } finally {
            setLoading(false);
        }
    };

    if (loadingProject) {
        return (
            <div className="min-h-screen bg-gray-50 flex items-center justify-center">
                <div className="text-muted-foreground">Loading project...</div>
            </div>
        );
    }

    if (projectError) {
        return (
            <div className="min-h-screen bg-gray-50 flex items-center justify-center">
                <div className="text-center space-y-4">
                    <p className="text-red-500">{projectError}</p>
                    <div className="space-x-2">
                        <Button onClick={fetchProject}>Retry</Button>
                        <Button variant="outline" onClick={() => navigate("/dashboard")}>
                            Back to Dashboard
                        </Button>
                    </div>
                </div>
            </div>
        );
    }

    if (!project) {
        return (
            <div className="min-h-screen bg-gray-50 flex items-center justify-center">
                <div className="text-muted-foreground">Project not found</div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-gray-50">
            <div className="border-b bg-white">
                <div className="max-w-7xl mx-auto px-8 py-6">
                    <div className="flex items-center gap-4 mb-4">
                        <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => navigate("/dashboard")}
                        >
                            <ArrowLeft className="h-4 w-4 mr-2" />
                            Back to Projects
                        </Button>
                    </div>
                    <div className="flex items-center justify-between">
                        <div>
                            <h1 className="text-3xl font-bold tracking-tight">
                                {project.projectName}
                            </h1>
                            <p className="text-muted-foreground mt-1">
                                {project.projectDescription}
                            </p>
                            <div className="flex gap-2 mt-3">
                                <span className="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold">
                                    {project.projectType}
                                </span>
                                {project.isMultiTenant && (
                                    <span className="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold">
                                        Multi-tenant
                                    </span>
                                )}
                                {project.isMultiLingual && (
                                    <span className="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold">
                                        Multi-lingual
                                    </span>
                                )}
                            </div>
                        </div>
                        <Dialog open={open} onOpenChange={setOpen}>
                            <DialogTrigger asChild>
                                <Button>
                                    <Plus className="mr-2 h-4 w-4" /> Add Entity
                                </Button>
                            </DialogTrigger>
                            <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
                                <DialogHeader>
                                    <DialogTitle>Create Entity</DialogTitle>
                                    <DialogDescription>
                                        Add a new entity to your project.
                                    </DialogDescription>
                                </DialogHeader>
                                <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                                    <div className="space-y-2">
                                        <Label htmlFor="entityName">Entity Name</Label>
                                        <Input
                                            id="entityName"
                                            placeholder="User"
                                            {...register("entityName")}
                                        />
                                        {errors.entityName && (
                                            <p className="text-sm text-red-500">
                                                {errors.entityName.message}
                                            </p>
                                        )}
                                    </div>

                                    <div className="space-y-2">
                                        <Label htmlFor="entityDescription">Description</Label>
                                        <Input
                                            id="entityDescription"
                                            placeholder="Brief description..."
                                            {...register("entityDescription")}
                                        />
                                    </div>

                                    <div className="space-y-2">
                                        <Label htmlFor="version">Version</Label>
                                        <Input
                                            id="version"
                                            placeholder="1.0"
                                            {...register("version")}
                                        />
                                    </div>

                                    <div className="grid grid-cols-2 gap-4">
                                        <div className="space-y-2">
                                            <Label>Database</Label>
                                            <Controller
                                                name="preferredDB"
                                                control={control}
                                                render={({ field }) => (
                                                    <Select
                                                        onValueChange={field.onChange}
                                                        defaultValue={field.value}
                                                    >
                                                        <SelectTrigger>
                                                            <SelectValue placeholder="Select database" />
                                                        </SelectTrigger>
                                                        <SelectContent>
                                                            {DB_OPTIONS.map((db) => (
                                                                <SelectItem key={db} value={db}>
                                                                    {db}
                                                                </SelectItem>
                                                            ))}
                                                        </SelectContent>
                                                    </Select>
                                                )}
                                            />
                                        </div>

                                        <div className="space-y-2">
                                            <Label>DB Interaction</Label>
                                            <Controller
                                                name="modeOfDBInteraction"
                                                control={control}
                                                render={({ field }) => (
                                                    <Select
                                                        onValueChange={field.onChange}
                                                        defaultValue={field.value}
                                                    >
                                                        <SelectTrigger>
                                                            <SelectValue placeholder="Select mode" />
                                                        </SelectTrigger>
                                                        <SelectContent>
                                                            {DB_INTERACTION_MODES.map((mode) => (
                                                                <SelectItem key={mode} value={mode}>
                                                                    {mode}
                                                                </SelectItem>
                                                            ))}
                                                        </SelectContent>
                                                    </Select>
                                                )}
                                            />
                                        </div>
                                    </div>

                                    <div className="space-y-4">
                                        <div className="flex items-center justify-between">
                                            <Label>RBAC</Label>
                                            <Controller
                                                name="implementsRBAC"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex items-center justify-between">
                                            <Label>Authentication Required</Label>
                                            <Controller
                                                name="isAuthenticationRequired"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex items-center justify-between">
                                            <Label>Audit Trail</Label>
                                            <Controller
                                                name="implementsAudit"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex items-center justify-between">
                                            <Label>Change Management</Label>
                                            <Controller
                                                name="implementsChangeManagement"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex items-center justify-between">
                                            <Label>Read Only</Label>
                                            <Controller
                                                name="isReadOnly"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex items-center justify-between">
                                            <Label>Independent Entity</Label>
                                            <Controller
                                                name="isIndependentEntity"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex items-center justify-between">
                                            <Label>Backend Only</Label>
                                            <Controller
                                                name="isBackendOnly"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>
                                    </div>

                                    <DialogFooter>
                                        <Button type="submit" disabled={loading}>
                                            {loading ? "Creating..." : "Create Entity"}
                                        </Button>
                                    </DialogFooter>
                                </form>
                            </DialogContent>
                        </Dialog>
                    </div>
                </div>
            </div>

            <div className="max-w-7xl mx-auto px-8 py-8">
                <Tabs defaultValue="entities" className="space-y-6">
                    <TabsList>
                        <TabsTrigger value="entities">Entities</TabsTrigger>
                        <TabsTrigger value="journeys">Journeys</TabsTrigger>
                    </TabsList>

                    <TabsContent value="entities" className="space-y-6">
                        <div className="mb-6">
                            <h2 className="text-xl font-semibold">Entities</h2>
                            <p className="text-sm text-muted-foreground">
                                Manage the entities in your project
                            </p>
                        </div>

                        {project.entities && project.entities.length > 0 ? (
                            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                                {project.entities.map((entity) => (
                                    <Card
                                        key={entity.uuid}
                                        className="hover:shadow-md transition-shadow cursor-pointer"
                                        onClick={() => navigate(`/projects/${id}/entities/${entity.uuid}`)}
                                    >
                                        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                                            <CardTitle className="text-lg font-medium">
                                                {entity.entityName}
                                            </CardTitle>
                                            <div className="flex gap-2" onClick={(e) => e.stopPropagation()}>
                                                <Button variant="ghost" size="sm">
                                                    <Edit className="h-4 w-4" />
                                                </Button>
                                                <Button variant="ghost" size="sm">
                                                    <Trash2 className="h-4 w-4 text-red-500" />
                                                </Button>
                                            </div>
                                        </CardHeader>
                                        <CardContent>
                                            <p className="text-sm text-muted-foreground mb-4">
                                                {entity.entityDescription || "No description"}
                                            </p>
                                            <div className="flex flex-wrap gap-1">
                                                <span className="inline-flex items-center rounded-full bg-blue-50 px-2 py-1 text-xs font-medium text-blue-700">
                                                    {entity.preferredDB}
                                                </span>
                                                <span className="inline-flex items-center rounded-full bg-purple-50 px-2 py-1 text-xs font-medium text-purple-700">
                                                    {entity.modeOfDBInteraction}
                                                </span>
                                                {entity.implementsRBAC && (
                                                    <span className="inline-flex items-center rounded-full bg-green-50 px-2 py-1 text-xs font-medium text-green-700">
                                                        RBAC
                                                    </span>
                                                )}
                                                {entity.isAuthenticationRequired && (
                                                    <span className="inline-flex items-center rounded-full bg-orange-50 px-2 py-1 text-xs font-medium text-orange-700">
                                                        Auth
                                                    </span>
                                                )}
                                                {entity.implementsAudit && (
                                                    <span className="inline-flex items-center rounded-full bg-yellow-50 px-2 py-1 text-xs font-medium text-yellow-700">
                                                        Audit
                                                    </span>
                                                )}
                                            </div>
                                        </CardContent>
                                    </Card>
                                ))}
                            </div>
                        ) : (
                            <Card className="p-12 text-center">
                                <div className="text-muted-foreground">
                                    <p className="mb-4">No entities yet</p>
                                    <Button onClick={() => setOpen(true)}>
                                        <Plus className="mr-2 h-4 w-4" /> Create your first entity
                                    </Button>
                                </div>
                            </Card>
                        )}
                    </TabsContent>

                    <TabsContent value="journeys">
                        <div className="mb-6">
                            <h2 className="text-xl font-semibold">Journeys</h2>
                            <p className="text-sm text-muted-foreground">
                                Manage the journeys and workflows for your project
                            </p>
                        </div>
                        <JourneyList projectId={id!} />
                    </TabsContent>
                </Tabs>
            </div>
        </div>
    );
}
