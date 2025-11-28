import { useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Plus, Folder, FileText } from "lucide-react";
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
import BlueprintForm from "@/components/BlueprintForm";
import type { BlueprintFormData } from "@/components/BlueprintForm";
import type { Blueprint } from "@/types/blueprint";
import { useNavigate } from "react-router-dom";

const projectSchema = z.object({
    projectName: z.string().min(2, "Project name is required"),
    projectDescription: z.string().optional(),
    projectType: z.string().default("Enterprise"),
    isMultiTenant: z.boolean().default(false),
    isMultiLingual: z.boolean().default(false),
});

type ProjectFormData = z.infer<typeof projectSchema>;

interface Project {
    uuid: string;
    projectName: string;
    projectDescription: string;
    projectType: string;
    isMultiTenant: boolean;
    isMultiLingual: boolean;
}



const PROJECT_TYPES = [
    "Enterprise",
    "Website",
    "MobileApp",
    "DesktopApp",
    "Microservice",
    "API",
    "SaaS",
    "ECommerce",
    "CRM",
    "CMS",
];

export default function DashboardPage() {
    const navigate = useNavigate();
    const [projects, setProjects] = useState<Project[]>([]);
    const [blueprints, setBlueprints] = useState<Blueprint[]>([]);
    const [open, setOpen] = useState(false);
    const [blueprintOpen, setBlueprintOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const {
        register,
        handleSubmit,
        reset,
        control,
        formState: { errors },
    } = useForm<ProjectFormData>({
        resolver: zodResolver(projectSchema),
        defaultValues: {
            projectName: "",
            projectDescription: "",
            projectType: "Enterprise",
            isMultiTenant: false,
            isMultiLingual: false,
        },
    });



    const fetchProjects = async () => {
        try {
            const res = await api.post<{ items: Project[] }>("/projects/get-by-filter", {
                pageNumber: 1,
                pageSize: 100,
            });
            if (res && res.items) {
                setProjects(res.items);
            }
        } catch (error) {
            console.error("Failed to fetch projects", error);
        }
    };

    const fetchBlueprints = async () => {
        try {
            const res = await api.post<{ items: Blueprint[] }>("/blueprints/get-by-filter", {
                pageNumber: 1,
                pageSize: 100,
            });
            if (res && res.items) {
                setBlueprints(res.items);
            }
        } catch (error) {
            console.error("Failed to fetch blueprints", error);
        }
    };

    useEffect(() => {
        fetchProjects();
        fetchBlueprints();
    }, []);

    const onSubmit = async (data: ProjectFormData) => {
        setLoading(true);
        try {
            await api.post("/projects/", data);
            setOpen(false);
            reset();
            fetchProjects();
        } catch (error) {
            console.error("Failed to create project", error);
        } finally {
            setLoading(false);
        }
    };

    const onBlueprintSubmit = async (data: BlueprintFormData) => {
        setLoading(true);
        try {
            await api.post("/blueprints/", data);
            setBlueprintOpen(false);
            fetchBlueprints();
        } catch (error) {
            console.error("Failed to create blueprint", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-gray-50 p-8">
            <div className="max-w-6xl mx-auto space-y-8">
                <div>
                    <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
                    <p className="text-muted-foreground">
                        Manage your enterprise projects and blueprints.
                    </p>
                </div>

                <Tabs defaultValue="projects" className="w-full">
                    <TabsList className="grid w-full grid-cols-2 max-w-md">
                        <TabsTrigger value="projects">Projects</TabsTrigger>
                        <TabsTrigger value="blueprints">Blueprints</TabsTrigger>
                    </TabsList>

                    <TabsContent value="projects" className="space-y-4 mt-6">
                        <div className="flex items-center justify-between">
                            <div>
                                <h2 className="text-2xl font-semibold">Projects</h2>
                                <p className="text-sm text-muted-foreground">
                                    View and manage your projects
                                </p>
                            </div>
                            <Dialog open={open} onOpenChange={setOpen}>
                                <DialogTrigger asChild>
                                    <Button>
                                        <Plus className="mr-2 h-4 w-4" /> New Project
                                    </Button>
                                </DialogTrigger>
                                <DialogContent className="sm:max-w-[500px]">
                                    <DialogHeader>
                                        <DialogTitle>Create Project</DialogTitle>
                                        <DialogDescription>
                                            Add a new project to your organization.
                                        </DialogDescription>
                                    </DialogHeader>
                                    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                                        <div className="space-y-2">
                                            <Label htmlFor="name">Project Name</Label>
                                            <Input
                                                id="name"
                                                placeholder="My Enterprise App"
                                                {...register("projectName")}
                                            />
                                            {errors.projectName && (
                                                <p className="text-sm text-red-500">
                                                    {errors.projectName.message}
                                                </p>
                                            )}
                                        </div>
                                        <div className="space-y-2">
                                            <Label htmlFor="description">Description</Label>
                                            <Input
                                                id="description"
                                                placeholder="Brief description..."
                                                {...register("projectDescription")}
                                            />
                                        </div>

                                        <div className="space-y-2">
                                            <Label>Project Type</Label>
                                            <Controller
                                                name="projectType"
                                                control={control}
                                                render={({ field }) => (
                                                    <Select
                                                        onValueChange={field.onChange}
                                                        defaultValue={field.value}
                                                    >
                                                        <SelectTrigger>
                                                            <SelectValue placeholder="Select project type" />
                                                        </SelectTrigger>
                                                        <SelectContent>
                                                            {PROJECT_TYPES.map((type) => (
                                                                <SelectItem key={type} value={type}>
                                                                    {type}
                                                                </SelectItem>
                                                            ))}
                                                        </SelectContent>
                                                    </Select>
                                                )}
                                            />
                                        </div>

                                        <div className="flex flex-row items-center justify-between rounded-lg border p-4">
                                            <div className="space-y-0.5">
                                                <Label className="text-base">Multi-tenant</Label>
                                                <div className="text-sm text-muted-foreground">
                                                    Enable multi-tenancy support
                                                </div>
                                            </div>
                                            <Controller
                                                name="isMultiTenant"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <div className="flex flex-row items-center justify-between rounded-lg border p-4">
                                            <div className="space-y-0.5">
                                                <Label className="text-base">Multi-lingual</Label>
                                                <div className="text-sm text-muted-foreground">
                                                    Enable multi-language support
                                                </div>
                                            </div>
                                            <Controller
                                                name="isMultiLingual"
                                                control={control}
                                                render={({ field }) => (
                                                    <Switch
                                                        checked={field.value}
                                                        onCheckedChange={field.onChange}
                                                    />
                                                )}
                                            />
                                        </div>

                                        <DialogFooter>
                                            <Button type="submit" disabled={loading}>
                                                {loading ? "Creating..." : "Create Project"}
                                            </Button>
                                        </DialogFooter>
                                    </form>
                                </DialogContent>
                            </Dialog>
                        </div>

                        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
                            {projects.map((project) => (
                                <Card
                                    key={project.uuid}
                                    className="cursor-pointer hover:shadow-md transition-shadow"
                                    onClick={() => navigate(`/projects/${project.uuid}`)}
                                >
                                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                                        <CardTitle className="text-sm font-medium">
                                            {project.projectType}
                                        </CardTitle>
                                        <Folder className="h-4 w-4 text-muted-foreground" />
                                    </CardHeader>
                                    <CardContent>
                                        <div className="text-2xl font-bold">{project.projectName}</div>
                                        <p className="text-xs text-muted-foreground mt-1">
                                            {project.projectDescription}
                                        </p>
                                        <div className="flex gap-2 mt-4">
                                            {project.isMultiTenant && (
                                                <span className="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 text-foreground">
                                                    Multi-tenant
                                                </span>
                                            )}
                                            {project.isMultiLingual && (
                                                <span className="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 text-foreground">
                                                    Multi-lingual
                                                </span>
                                            )}
                                        </div>
                                    </CardContent>
                                </Card>
                            ))}
                        </div>
                    </TabsContent>

                    <TabsContent value="blueprints" className="space-y-4 mt-6">
                        <div className="flex items-center justify-between">
                            <div>
                                <h2 className="text-2xl font-semibold">Blueprints</h2>
                                <p className="text-sm text-muted-foreground">
                                    Manage your blueprint templates
                                </p>
                            </div>
                            <Dialog open={blueprintOpen} onOpenChange={setBlueprintOpen}>
                                <DialogTrigger asChild>
                                    <Button>
                                        <Plus className="mr-2 h-4 w-4" /> New Blueprint
                                    </Button>
                                </DialogTrigger>
                                <DialogContent className="sm:max-w-[700px] max-h-[80vh] overflow-y-auto">
                                    <DialogHeader>
                                        <DialogTitle>Create Blueprint</DialogTitle>
                                        <DialogDescription>
                                            Add a new blueprint template with functionalities.
                                        </DialogDescription>
                                    </DialogHeader>
                                    <BlueprintForm
                                        onSubmit={onBlueprintSubmit}
                                        onCancel={() => setBlueprintOpen(false)}
                                        isLoading={loading}
                                    />
                                </DialogContent>
                            </Dialog>
                        </div>

                        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
                            {blueprints.map((blueprint) => (
                                <Card
                                    key={blueprint.uuid}
                                    className="cursor-pointer hover:shadow-md transition-shadow"
                                    onClick={() => navigate(`/blueprints/${blueprint.uuid}`)}
                                >
                                    <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                                        <CardTitle className="text-sm font-medium">
                                            {blueprint.type}
                                        </CardTitle>
                                        <FileText className="h-4 w-4 text-muted-foreground" />
                                    </CardHeader>
                                    <CardContent>
                                        <div className="text-2xl font-bold">{blueprint.standardName}</div>
                                        <p className="text-xs text-muted-foreground mt-1">
                                            {blueprint.description}
                                        </p>
                                        {blueprint.functionalities && blueprint.functionalities.length > 0 && (
                                            <div className="flex gap-2 mt-4">
                                                <span className="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 text-foreground">
                                                    {blueprint.functionalities.length} {blueprint.functionalities.length === 1 ? 'functionality' : 'functionalities'}
                                                </span>
                                            </div>
                                        )}
                                    </CardContent>
                                </Card>
                            ))}
                        </div>
                    </TabsContent>
                </Tabs>
            </div>
        </div>
    );
}
