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
    const [open, setOpen] = useState(false);
    const [loading, setLoading] = useState(false);

    const {
        register,
        handleSubmit,
        reset,
        control,
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
        },
    });

    const isEnum = watch("isEnum");
    const isDerived = watch("isDerived");
    const isCollection = watch("isCollection");

    const fetchEntity = async () => {
        try {
            const res = await api.get<Entity>(`/entities/${entityId}`);
            setEntity(res);
        } catch (error) {
            console.error("Failed to fetch entity", error);
        }
    };

    useEffect(() => {
        if (entityId) {
            fetchEntity();
        }
    }, [entityId]);

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

    if (!entity) {
        return (
            <div className="min-h-screen bg-gray-50 flex items-center justify-center">
                <div className="text-muted-foreground">Loading...</div>
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

                                    <div className="grid grid-cols-2 gap-4">
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
                                    </div>

                                    <div className="space-y-2">
                                        <Label htmlFor="sampleData">Sample Data</Label>
                                        <Input
                                            id="sampleData"
                                            placeholder="john.doe@example.com"
                                            {...register("sampleData")}
                                        />
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

                                    <div className="border-t pt-4">
                                        <h4 className="font-medium mb-4">Advanced Options</h4>
                                        <div className="space-y-4">
                                            <div className="flex items-center justify-between">
                                                <Label>Is Enum</Label>
                                                <Controller
                                                    name="isEnum"
                                                    control={control}
                                                    render={({ field }) => (
                                                        <Switch
                                                            checked={field.value}
                                                            onCheckedChange={field.onChange}
                                                        />
                                                    )}
                                                />
                                            </div>

                                            {isEnum && (
                                                <div className="space-y-2">
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

                                            <div className="flex items-center justify-between">
                                                <Label>Is Collection</Label>
                                                <Controller
                                                    name="isCollection"
                                                    control={control}
                                                    render={({ field }) => (
                                                        <Switch
                                                            checked={field.value}
                                                            onCheckedChange={field.onChange}
                                                        />
                                                    )}
                                                />
                                            </div>

                                            {isCollection && (
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
                                            )}

                                            <div className="flex items-center justify-between">
                                                <Label>Is Derived</Label>
                                                <Controller
                                                    name="isDerived"
                                                    control={control}
                                                    render={({ field }) => (
                                                        <Switch
                                                            checked={field.value}
                                                            onCheckedChange={field.onChange}
                                                        />
                                                    )}
                                                />
                                            </div>

                                            {isDerived && (
                                                <>
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
                                                                            <SelectItem
                                                                                key={type}
                                                                                value={type}
                                                                            >
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
                                                </>
                                            )}
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
                                            <Button variant="ghost" size="sm">
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
                                            <span className="inline-flex items-center rounded-full bg-gray-50 px-2.5 py-0.5 text-xs font-medium text-gray-700">
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
        </div>
    );
}
