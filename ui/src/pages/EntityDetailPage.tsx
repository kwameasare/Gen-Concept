import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Plus, ArrowLeft, Trash2, Edit } from "lucide-react";
import { api } from "@/lib/api";
import { DeleteConfirmDialog } from "@/components/DeleteConfirmDialog";
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

const fieldSchema = z.object({
    fieldName: z.string().min(1, "Field name is required"),
    displayName: z.string().min(1, "Display name is required"),
    fieldDescription: z.string().optional(),
    fieldType: z.string().default("String"),
    isMandatory: z.boolean().default(false),
    isUnique: z.boolean().default(false),
    isReadOnly: z.boolean().default(false),
    isSensitive: z.boolean().default(false),
    isEditable: z.boolean().default(true),
    isDerived: z.boolean().default(false),
    isCollection: z.boolean().default(false),
    collectionType: z.string().default("None"),
    isEnum: z.boolean().default(false),
    enumValues: z.string().optional(),
    derivativeType: z.string().default("None"),
    derivativeExpression: z.string().optional(),
    isBackendOnly: z.boolean().default(false),
    displayStatus: z.string().default("Show"),
    sampleData: z.string().optional(),
    collectionItemType: z.string().default("Primitive"),
    collectionEntity: z.string().optional(),
});

type FieldFormData = z.infer<typeof fieldSchema>;

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
    entityFields: EntityField[];
}

interface Project {
    uuid: string;
    projectName: string;
    entities: Entity[];
}

interface EntityField {
    uuid: string;
    fieldName: string;
    displayName: string;
    fieldDescription: string;
    fieldType: string;
    isMandatory: boolean;
    isUnique: boolean;
    isReadOnly: boolean;
    isSensitive: boolean;
    isEditable: boolean;
    isDerived: boolean;
    isCollection: boolean;
    collectionType: string;
    isEnum: boolean;
    enumValues: string[];
    derivativeType: string;
    derivativeExpression: string;
    isBackendOnly: boolean;
    displayStatus: string;
    sampleData: string;
    collectionItemType: string;
    collectionEntity: string;
}

const FIELD_TYPES = [
    "String",
    "Integer",
    "Float",
    "Boolean",
    "Date",
    "DateTime",
    "Time",
    "Text",
    "JSON",
    "UUID",
    "Email",
    "URL",
];

const COLLECTION_TYPES = ["None", "List", "Set", "Map"];
const DERIVATIVE_TYPES = ["None", "Computed", "Concatenated", "Lookup"];
const DISPLAY_STATUS = ["Show", "Detail", "Hide"];

export default function EntityDetailPage() {
    const { projectId, entityId } = useParams<{ projectId: string; entityId: string }>();
    const navigate = useNavigate();
    const [entity, setEntity] = useState<Entity | null>(null);
    const [project, setProject] = useState<Project | null>(null);
    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);
    const [createEntityOpen, setCreateEntityOpen] = useState(false);
    const [newEntityName, setNewEntityName] = useState("");
    const [deleteEntityDialog, setDeleteEntityDialog] = useState(false);
    const [deleteFieldDialog, setDeleteFieldDialog] = useState<{ open: boolean; field: EntityField | null }>({
        open: false,
        field: null,
    });

    const {
        register,
        handleSubmit,
        reset,
        control,
        setValue,
        watch,
        formState: { errors },
    } = useForm<FieldFormData>({
        resolver: zodResolver(fieldSchema),
        defaultValues: {
            fieldName: "",
            displayName: "",
            fieldDescription: "",
            fieldType: "String",
            isMandatory: false,
            isUnique: false,
            isReadOnly: false,
            isSensitive: false,
            isEditable: true,
            isDerived: false,
            isCollection: false,
            collectionType: "None",
            isEnum: false,
            enumValues: "",
            derivativeType: "None",
            derivativeExpression: "",
            isBackendOnly: false,
            displayStatus: "Show",
            sampleData: "",
            collectionItemType: "Primitive",
            collectionEntity: "",
        },
    });

    const isEnum = watch("isEnum");
    const isDerived = watch("isDerived");
    const isCollection = watch("isCollection");
    const collectionItemType = watch("collectionItemType");

    // Mutual exclusivity logic
    useEffect(() => {
        const subscription = watch((value, { name }) => {
            if (name === "isEnum" && value.isEnum) {
                setValue("isCollection", false);
                setValue("isDerived", false);
            } else if (name === "isCollection" && value.isCollection) {
                setValue("isEnum", false);
                setValue("isDerived", false);
            } else if (name === "isDerived" && value.isDerived) {
                setValue("isEnum", false);
                setValue("isCollection", false);
            }
        });
        return () => subscription.unsubscribe();
    }, [watch]);

    const fetchProject = async () => {
        try {
            const res = await api.get<Project>(`/projects/${projectId}`);
            setProject(res);
        } catch (error) {
            console.error("Failed to fetch project", error);
        }
    };

    const fetchEntity = async () => {
        try {
            const res = await api.get<Entity>(`/entities/${entityId}`);
            setEntity(res);
        } catch (error) {
            console.error("Failed to fetch entity", error);
        }
    };

    useEffect(() => {
        if (projectId) {
            fetchProject();
        }
    }, [projectId]);

    useEffect(() => {
        if (entityId) {
            fetchEntity();
        }
    }, [entityId]);

    const handleCreateEntity = async () => {
        if (!newEntityName) return;
        try {
            const payload = {
                entityName: newEntityName,
                entityDescription: "Ad-hoc entity created from field definition",
                projectId: projectId,
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
            };
            await api.post("/entities/", payload);
            setCreateEntityOpen(false);
            setNewEntityName("");
            fetchProject(); // Refresh project to get new entity list
        } catch (error) {
            console.error("Failed to create entity", error);
        }
    };

    const onSubmit = async (data: FieldFormData) => {
        setLoading(true);
        try {
            const payload = {
                ...data,
                enumValues: data.enumValues ? data.enumValues.split(",").map((v) => v.trim()) : [],
            };
            await api.post(`/entity-fields/`, {
                ...payload,
                entityId,
            });
            setOpen(false);
            reset();
            fetchEntity();
        } catch (error) {
            console.error("Failed to create field", error);
        } finally {
            setLoading(false);
        }
    };

    const handleDeleteEntity = async () => {
        try {
            await api.delete(`/entities/${entityId}`);
            navigate(`/projects/${projectId}`);
        } catch (error) {
            console.error("Failed to delete entity", error);
        }
    };

    const handleDeleteField = async () => {
        if (!deleteFieldDialog.field) return;
        try {
            await api.delete(`/entity-fields/${deleteFieldDialog.field.uuid}`);
            fetchEntity();
        } catch (error) {
            console.error("Failed to delete field", error);
        }
    };

    if (!entity) {
        return (
            <div className="min-h-screen bg-background flex items-center justify-center">
                <div className="text-muted-foreground">Loading...</div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-background">
            <div className="border-b bg-card">
                <div className="max-w-7xl mx-auto px-8 py-6">
                    <div className="flex items-center gap-4 mb-4">
                        <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => navigate(`/projects/${projectId}`)}
                        >
                            <ArrowLeft className="h-4 w-4 mr-2" />
                            Back to Project
                        </Button>
                    </div>
                    <div className="flex items-center justify-between">
                        <div>
                            <h1 className="text-3xl font-bold tracking-tight">
                                {entity.entityName}
                            </h1>
                            <p className="text-muted-foreground mt-1">
                                {entity.entityDescription || "No description"}
                            </p>
                            <div className="flex gap-2 mt-3">
                                <span className="inline-flex items-center rounded-full bg-blue-50 px-2.5 py-0.5 text-xs font-medium text-blue-700">
                                    {entity.preferredDB}
                                </span>
                                <span className="inline-flex items-center rounded-full bg-purple-50 px-2.5 py-0.5 text-xs font-medium text-purple-700">
                                    {entity.modeOfDBInteraction}
                                </span>
                                {entity.implementsRBAC && (
                                    <span className="inline-flex items-center rounded-full bg-green-50 px-2.5 py-0.5 text-xs font-medium text-green-700">
                                        RBAC
                                    </span>
                                )}
                            </div>
                        </div>
                        <div className="flex gap-2">
                            <Dialog open={open} onOpenChange={setOpen}>
                                <DialogTrigger asChild>
                                    <Button>
                                        <Plus className="mr-2 h-4 w-4" /> Add Field
                                    </Button>
                                </DialogTrigger>
                                <DialogContent className="sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
                                    <DialogHeader>
                                        <DialogTitle>Create Field</DialogTitle>
                                        <DialogDescription>
                                            Add a new field to {entity.entityName}
                                        </DialogDescription>
                                    </DialogHeader>
                                    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                                        <div className="grid grid-cols-2 gap-4">
                                            <div className="space-y-2">
                                                <Label htmlFor="fieldName">Field Name</Label>
                                                <Input
                                                    id="fieldName"
                                                    placeholder="email"
                                                    {...register("fieldName")}
                                                />
                                                {errors.fieldName && (
                                                    <p className="text-sm text-red-500">
                                                        {errors.fieldName.message}
                                                    </p>
                                                )}
                                            </div>

                                            <div className="space-y-2">
                                                <Label htmlFor="displayName">Display Name</Label>
                                                <Input
                                                    id="displayName"
                                                    placeholder="Email Address"
                                                    {...register("displayName")}
                                                />
                                                {errors.displayName && (
                                                    <p className="text-sm text-red-500">
                                                        {errors.displayName.message}
                                                    </p>
                                                )}
                                            </div>
                                        </div>

                                        <div className="space-y-2">
                                            <Label htmlFor="fieldDescription">Description</Label>
                                            <Input
                                                id="fieldDescription"
                                                placeholder="User's email address"
                                                {...register("fieldDescription")}
                                            />
                                        </div>

                                        <div className="border-t pt-4">
                                            <h4 className="font-medium mb-4">Data Structure</h4>
                                            <div className="space-y-4">
                                                <div className="flex gap-4">
                                                    <div className="flex items-center space-x-2">
                                                        <Controller
                                                            name="isEnum"
                                                            control={control}
                                                            render={({ field }) => (
                                                                <Switch
                                                                    id="isEnum"
                                                                    checked={field.value}
                                                                    onCheckedChange={field.onChange}
                                                                />
                                                            )}
                                                        />
                                                        <Label htmlFor="isEnum">Is Enum</Label>
                                                    </div>
                                                    <div className="flex items-center space-x-2">
                                                        <Controller
                                                            name="isCollection"
                                                            control={control}
                                                            render={({ field }) => (
                                                                <Switch
                                                                    id="isCollection"
                                                                    checked={field.value}
                                                                    onCheckedChange={field.onChange}
                                                                />
                                                            )}
                                                        />
                                                        <Label htmlFor="isCollection">Is Collection</Label>
                                                    </div>
                                                    <div className="flex items-center space-x-2">
                                                        <Controller
                                                            name="isDerived"
                                                            control={control}
                                                            render={({ field }) => (
                                                                <Switch
                                                                    id="isDerived"
                                                                    checked={field.value}
                                                                    onCheckedChange={field.onChange}
                                                                />
                                                            )}
                                                        />
                                                        <Label htmlFor="isDerived">Is Derived</Label>
                                                    </div>
                                                </div>

                                                {/* Enum Configuration */}
                                                {isEnum && (
                                                    <div className="space-y-2 p-4 bg-muted rounded-md">
                                                        <Label htmlFor="enumValues">
                                                            Enum Values (comma-separated)
                                                        </Label>
                                                        <Input
                                                            id="enumValues"
                                                            placeholder="active, inactive, pending"
                                                            {...register("enumValues")}
                                                        />
                                                    </div>
                                                )}

                                                {/* Collection Configuration */}
                                                {isCollection && (
                                                    <div className="space-y-4 p-4 bg-muted rounded-md">
                                                        <div className="grid grid-cols-2 gap-4">
                                                            <div className="space-y-2">
                                                                <Label>Collection Type</Label>
                                                                <Controller
                                                                    name="collectionType"
                                                                    control={control}
                                                                    render={({ field }) => (
                                                                        <Select
                                                                            onValueChange={field.onChange}
                                                                            defaultValue={field.value}
                                                                        >
                                                                            <SelectTrigger>
                                                                                <SelectValue placeholder="Select type" />
                                                                            </SelectTrigger>
                                                                            <SelectContent>
                                                                                {COLLECTION_TYPES.map((type) => (
                                                                                    <SelectItem key={type} value={type}>
                                                                                        {type}
                                                                                    </SelectItem>
                                                                                ))}
                                                                            </SelectContent>
                                                                        </Select>
                                                                    )}
                                                                />
                                                            </div>
                                                            <div className="space-y-2">
                                                                <Label>Item Type</Label>
                                                                <Controller
                                                                    name="collectionItemType"
                                                                    control={control}
                                                                    render={({ field }) => (
                                                                        <Select
                                                                            onValueChange={field.onChange}
                                                                            defaultValue={field.value}
                                                                        >
                                                                            <SelectTrigger>
                                                                                <SelectValue placeholder="Select item type" />
                                                                            </SelectTrigger>
                                                                            <SelectContent>
                                                                                <SelectItem value="Primitive">Primitive</SelectItem>
                                                                                <SelectItem value="Entity">Entity</SelectItem>
                                                                            </SelectContent>
                                                                        </Select>
                                                                    )}
                                                                />
                                                            </div>
                                                        </div>

                                                        {collectionItemType === "Entity" && (
                                                            <div className="space-y-2">
                                                                <Label>Select Entity</Label>
                                                                <div className="flex gap-2">
                                                                    <Controller
                                                                        name="collectionEntity"
                                                                        control={control}
                                                                        render={({ field }) => (
                                                                            <Select
                                                                                onValueChange={field.onChange}
                                                                                defaultValue={field.value}
                                                                            >
                                                                                <SelectTrigger className="flex-1">
                                                                                    <SelectValue placeholder="Select entity" />
                                                                                </SelectTrigger>
                                                                                <SelectContent>
                                                                                    {project?.entities?.map((e) => (
                                                                                        <SelectItem key={e.uuid} value={e.uuid}>
                                                                                            {e.entityName}
                                                                                        </SelectItem>
                                                                                    ))}
                                                                                </SelectContent>
                                                                            </Select>
                                                                        )}
                                                                    />
                                                                    <Dialog open={createEntityOpen} onOpenChange={setCreateEntityOpen}>
                                                                        <DialogTrigger asChild>
                                                                            <Button variant="outline" size="icon">
                                                                                <Plus className="h-4 w-4" />
                                                                            </Button>
                                                                        </DialogTrigger>
                                                                        <DialogContent>
                                                                            <DialogHeader>
                                                                                <DialogTitle>Create New Entity</DialogTitle>
                                                                                <DialogDescription>
                                                                                    Create a new entity to use in this collection.
                                                                                </DialogDescription>
                                                                            </DialogHeader>
                                                                            <div className="space-y-4 py-4">
                                                                                <div className="space-y-2">
                                                                                    <Label>Entity Name</Label>
                                                                                    <Input
                                                                                        value={newEntityName}
                                                                                        onChange={(e) => setNewEntityName(e.target.value)}
                                                                                        placeholder="e.g., OrderItem"
                                                                                    />
                                                                                </div>
                                                                            </div>
                                                                            <DialogFooter>
                                                                                <Button onClick={handleCreateEntity}>Create Entity</Button>
                                                                            </DialogFooter>
                                                                        </DialogContent>
                                                                    </Dialog>
                                                                </div>
                                                            </div>
                                                        )}
                                                    </div>
                                                )}

                                                {/* Derived Configuration */}
                                                {isDerived && (
                                                    <div className="space-y-4 p-4 bg-muted rounded-md">
                                                        <div className="space-y-2">
                                                            <Label>Derivative Type</Label>
                                                            <Controller
                                                                name="derivativeType"
                                                                control={control}
                                                                render={({ field }) => (
                                                                    <Select
                                                                        onValueChange={field.onChange}
                                                                        defaultValue={field.value}
                                                                    >
                                                                        <SelectTrigger>
                                                                            <SelectValue placeholder="Select type" />
                                                                        </SelectTrigger>
                                                                        <SelectContent>
                                                                            {DERIVATIVE_TYPES.map((type) => (
                                                                                <SelectItem key={type} value={type}>
                                                                                    {type}
                                                                                </SelectItem>
                                                                            ))}
                                                                        </SelectContent>
                                                                    </Select>
                                                                )}
                                                            />
                                                        </div>
                                                        <div className="space-y-2">
                                                            <Label htmlFor="derivativeExpression">
                                                                Derivative Expression
                                                            </Label>
                                                            <Input
                                                                id="derivativeExpression"
                                                                placeholder="firstName + ' ' + lastName"
                                                                {...register("derivativeExpression")}
                                                            />
                                                        </div>
                                                    </div>
                                                )}
                                            </div>
                                        </div>

                                        {/* Field Type - Only show if not Entity Collection */}
                                        {(!isCollection || collectionItemType === "Primitive") && (
                                            <div className="space-y-2">
                                                <Label>Field Type</Label>
                                                <Controller
                                                    name="fieldType"
                                                    control={control}
                                                    render={({ field }) => (
                                                        <Select
                                                            onValueChange={field.onChange}
                                                            defaultValue={field.value}
                                                        >
                                                            <SelectTrigger>
                                                                <SelectValue placeholder="Select type" />
                                                            </SelectTrigger>
                                                            <SelectContent>
                                                                {FIELD_TYPES.map((type) => (
                                                                    <SelectItem key={type} value={type}>
                                                                        {type}
                                                                    </SelectItem>
                                                                ))}
                                                            </SelectContent>
                                                        </Select>
                                                    )}
                                                />
                                            </div>
                                        )}

                                        <div className="grid grid-cols-2 gap-4">
                                            <div className="space-y-2">
                                                <Label>Display Status</Label>
                                                <Controller
                                                    name="displayStatus"
                                                    control={control}
                                                    render={({ field }) => (
                                                        <Select
                                                            onValueChange={field.onChange}
                                                            defaultValue={field.value}
                                                        >
                                                            <SelectTrigger>
                                                                <SelectValue placeholder="Select status" />
                                                            </SelectTrigger>
                                                            <SelectContent>
                                                                {DISPLAY_STATUS.map((status) => (
                                                                    <SelectItem key={status} value={status}>
                                                                        {status}
                                                                    </SelectItem>
                                                                ))}
                                                            </SelectContent>
                                                        </Select>
                                                    )}
                                                />
                                            </div>
                                            <div className="space-y-2">
                                                <Label htmlFor="sampleData">Sample Data</Label>
                                                <Input
                                                    id="sampleData"
                                                    placeholder="john.doe@example.com"
                                                    {...register("sampleData")}
                                                />
                                            </div>
                                        </div>

                                        <div className="border-t pt-4">
                                            <h4 className="font-medium mb-4">Field Properties</h4>
                                            <div className="grid grid-cols-2 gap-4">
                                                <div className="flex items-center justify-between">
                                                    <Label>Mandatory</Label>
                                                    <Controller
                                                        name="isMandatory"
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
                                                    <Label>Unique</Label>
                                                    <Controller
                                                        name="isUnique"
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
                                                    <Label>Sensitive</Label>
                                                    <Controller
                                                        name="isSensitive"
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
                                                    <Label>Editable</Label>
                                                    <Controller
                                                        name="isEditable"
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
                                        </div>

                                        <DialogFooter>
                                            <Button type="submit" disabled={loading}>
                                                {loading ? "Creating..." : "Create Field"}
                                            </Button>
                                        </DialogFooter>
                                    </form>
                                </DialogContent>
                            </Dialog>
                            <Button
                                variant="destructive"
                                onClick={() => setDeleteEntityDialog(true)}
                            >
                                <Trash2 className="mr-2 h-4 w-4" />
                                Delete Entity
                            </Button>
                        </div>
                    </div>
                </div>

                <div className="max-w-7xl mx-auto px-8 py-8">
                    <div className="mb-6">
                        <h2 className="text-xl font-semibold">Fields</h2>
                        <p className="text-sm text-muted-foreground">
                            Manage the fields in this entity
                        </p>
                    </div>

                    {entity.entityFields && entity.entityFields.length > 0 ? (
                        <div className="space-y-3">
                            {entity.entityFields.map((field) => (
                                <Card key={field.uuid}>
                                    <CardHeader className="pb-3">
                                        <div className="flex items-center justify-between">
                                            <div className="flex items-center gap-3">
                                                <CardTitle className="text-base font-medium">
                                                    {field.displayName}
                                                </CardTitle>
                                                <code className="text-sm text-muted-foreground bg-gray-100 px-2 py-0.5 rounded">
                                                    {field.fieldName}
                                                </code>
                                            </div>
                                            <div className="flex gap-2">
                                                <Button variant="ghost" size="sm">
                                                    <Edit className="h-4 w-4" />
                                                </Button>
                                                <Button
                                                    variant="ghost"
                                                    size="sm"
                                                    onClick={() => setDeleteFieldDialog({ open: true, field })}
                                                >
                                                    <Trash2 className="h-4 w-4 text-red-500" />
                                                </Button>
                                            </div>
                                        </div>
                                    </CardHeader>
                                    <CardContent>
                                        <p className="text-sm text-muted-foreground mb-3">
                                            {field.fieldDescription || "No description"}
                                        </p>
                                        <div className="flex flex-wrap gap-2">
                                            <span className="inline-flex items-center rounded-full bg-indigo-50 px-2.5 py-0.5 text-xs font-medium text-indigo-700">
                                                {field.fieldType}
                                            </span>
                                            {field.isMandatory && (
                                                <span className="inline-flex items-center rounded-full bg-red-50 px-2.5 py-0.5 text-xs font-medium text-red-700">
                                                    Required
                                                </span>
                                            )}
                                            {field.isUnique && (
                                                <span className="inline-flex items-center rounded-full bg-blue-50 px-2.5 py-0.5 text-xs font-medium text-blue-700">
                                                    Unique
                                                </span>
                                            )}
                                            {field.isSensitive && (
                                                <span className="inline-flex items-center rounded-full bg-orange-50 px-2.5 py-0.5 text-xs font-medium text-orange-700">
                                                    Sensitive
                                                </span>
                                            )}
                                            {field.isReadOnly && (
                                                <span className="inline-flex items-center rounded-full bg-muted px-2.5 py-0.5 text-xs font-medium text-foreground">
                                                    Read-only
                                                </span>
                                            )}
                                            {field.isEnum && (
                                                <span className="inline-flex items-center rounded-full bg-purple-50 px-2.5 py-0.5 text-xs font-medium text-purple-700">
                                                    Enum
                                                </span>
                                            )}
                                            {field.isCollection && (
                                                <span className="inline-flex items-center rounded-full bg-green-50 px-2.5 py-0.5 text-xs font-medium text-green-700">
                                                    {field.collectionType}
                                                </span>
                                            )}
                                            {field.isDerived && (
                                                <span className="inline-flex items-center rounded-full bg-yellow-50 px-2.5 py-0.5 text-xs font-medium text-yellow-700">
                                                    Derived
                                                </span>
                                            )}
                                            <span className="inline-flex items-center rounded-full bg-slate-50 px-2.5 py-0.5 text-xs font-medium text-slate-700">
                                                {field.displayStatus}
                                            </span>
                                        </div>
                                        {field.sampleData && (
                                            <div className="mt-3 text-sm">
                                                <span className="text-muted-foreground">Sample: </span>
                                                <code className="bg-gray-100 px-2 py-0.5 rounded">
                                                    {field.sampleData}
                                                </code>
                                            </div>
                                        )}
                                    </CardContent>
                                </Card>
                            ))}
                        </div>
                    ) : (
                        <Card className="p-12 text-center">
                            <div className="text-muted-foreground">
                                <p className="mb-4">No fields yet</p>
                                <Button onClick={() => setOpen(true)}>
                                    <Plus className="mr-2 h-4 w-4" /> Create your first field
                                </Button>
                            </div>
                        </Card>
                    )}
                </div>

                {/* Delete Confirmation Dialogs */}
                <DeleteConfirmDialog
                    open={deleteEntityDialog}
                    onOpenChange={setDeleteEntityDialog}
                    onConfirm={handleDeleteEntity}
                    title="Delete Entity"
                    description="Are you sure you want to delete this entity? All related fields and validations will also be deleted. This action cannot be undone."
                    itemName={entity?.entityName}
                />

                <DeleteConfirmDialog
                    open={deleteFieldDialog.open}
                    onOpenChange={(open) => setDeleteFieldDialog({ open, field: null })}
                    onConfirm={handleDeleteField}
                    title="Delete Field"
                    description="Are you sure you want to delete this field? This action cannot be undone."
                    itemName={deleteFieldDialog.field?.fieldName}
                />
            </div>
        </div>
    );
}
