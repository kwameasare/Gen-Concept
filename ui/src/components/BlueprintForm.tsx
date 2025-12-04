import { useEffect, useState } from "react";
import { useForm, useFieldArray, Controller, useWatch } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Textarea } from "@/components/ui/textarea";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Plus, Trash2, ChevronDown, ChevronUp } from "lucide-react";
import type { Blueprint, Library } from "@/types/blueprint";
import { api } from "@/lib/api";
import {
    FUNCTIONALITY_CATEGORIES,
    FUNCTIONALITY_TYPES,
    getProvidersForType,
} from "@/constants/functionality";

const functionalOperationSchema = z.object({
    uuid: z.string().optional(),
    name: z.string().min(1, "Operation name is required"),
    description: z.string().optional(),
});

const functionalitySchema = z.object({
    uuid: z.string().optional(),
    category: z.string().min(1, "Category is required"),
    type: z.string().min(1, "Type is required"),
    provider: z.string().optional(),
    implementsGenerics: z.boolean().default(false),
    filePathsCSV: z.string().optional(),
    operations: z.array(functionalOperationSchema).default([]),
});

const libraryReferenceSchema = z.object({
    uuid: z.string(),
    standardName: z.string(),
    version: z.string(),
    namespace: z.string().optional(),
    repositoryURL: z.string().optional(),
    exposedFunctionalities: z.array(z.any()).optional(),
});

const blueprintSchema = z.object({
    standardName: z.string().min(2, "Blueprint name is required"),
    type: z.string().min(1, "Type is required"),
    description: z.string().optional(),
    functionalities: z.array(functionalitySchema).default([]),
    libraries: z.array(libraryReferenceSchema).default([]),
});

export type BlueprintFormData = z.infer<typeof blueprintSchema>;

interface BlueprintFormProps {
    initialData?: Blueprint;
    onSubmit: (data: BlueprintFormData) => void;
    onCancel?: () => void;
    isLoading?: boolean;
}

export default function BlueprintForm({
    initialData,
    onSubmit,
    onCancel,
    isLoading = false,
}: BlueprintFormProps) {
    const {
        register,
        control,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<BlueprintFormData>({
        resolver: zodResolver(blueprintSchema),
        defaultValues: initialData || {
            standardName: "",
            type: "",
            description: "",
            functionalities: [],
            libraries: [],
        },
    });

    const [availableLibraries, setAvailableLibraries] = useState<Library[]>([]);

    useEffect(() => {
        fetchLibraries();
    }, []);

    const fetchLibraries = async () => {
        try {
            const response: any = await api.post("/libraries/get-by-filter", {
                pageNumber: 1,
                pageSize: 100,
            });
            setAvailableLibraries(response.items || []);
        } catch (error) {
            console.error("Error fetching libraries:", error);
        }
    };

    const {
        fields: functionalities,
        append: appendFunctionality,
        remove: removeFunctionality,
    } = useFieldArray({
        control,
        name: "functionalities",
    });

    const {
        fields: libraries,
        append: appendLibrary,
        remove: removeLibrary,
    } = useFieldArray({
        control,
        name: "libraries",
    });

    const [expandedFunctionalities, setExpandedFunctionalities] = useState<Set<number>>(
        new Set(functionalities.map((_, index) => index))
    );

    useEffect(() => {
        if (initialData) {
            reset(initialData);
        }
    }, [initialData, reset]);

    const toggleFunctionality = (index: number) => {
        setExpandedFunctionalities((prev) => {
            const newSet = new Set(prev);
            if (newSet.has(index)) {
                newSet.delete(index);
            } else {
                newSet.add(index);
            }
            return newSet;
        });
    };

    const addFunctionality = () => {
        appendFunctionality({
            category: "",
            type: "",
            provider: "",
            implementsGenerics: false,
            filePathsCSV: "",
            operations: [],
        });
        setExpandedFunctionalities((prev) => new Set([...prev, functionalities.length]));
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
            {/* Basic Fields */}
            <div className="space-y-4">
                <div className="space-y-2">
                    <Label htmlFor="standardName">Blueprint Name *</Label>
                    <Input
                        id="standardName"
                        placeholder="e.g., User Management"
                        {...register("standardName")}
                    />
                    {errors.standardName && (
                        <p className="text-sm text-red-500">{errors.standardName.message}</p>
                    )}
                </div>

                <div className="space-y-2">
                    <Label htmlFor="type">Type *</Label>
                    <Input
                        id="type"
                        placeholder="e.g., Authentication, CRUD"
                        {...register("type")}
                    />
                    {errors.type && (
                        <p className="text-sm text-red-500">{errors.type.message}</p>
                    )}
                </div>

                <div className="space-y-2">
                    <Label htmlFor="description">Description</Label>
                    <Textarea
                        id="description"
                        placeholder="Describe this blueprint..."
                        rows={3}
                        {...register("description")}
                    />
                </div>
            </div>

            {/* Functionalities Section */}
            <div className="space-y-4">
                <div className="flex items-center justify-between">
                    <h3 className="text-lg font-semibold">Functionalities</h3>
                    <Button
                        type="button"
                        variant="outline"
                        size="sm"
                        onClick={addFunctionality}
                    >
                        <Plus className="mr-2 h-4 w-4" />
                        Add Functionality
                    </Button>
                </div>

                {functionalities.map((functionality, funcIndex) => (
                    <FunctionalityCard
                        key={functionality.id}
                        funcIndex={funcIndex}
                        register={register}
                        control={control}
                        errors={errors}
                        isExpanded={expandedFunctionalities.has(funcIndex)}
                        onToggle={() => toggleFunctionality(funcIndex)}
                        onRemove={() => removeFunctionality(funcIndex)}
                    />
                ))}
            </div>

            {/* Libraries Section */}
            <div className="space-y-4">
                <div className="flex items-center justify-between">
                    <h3 className="text-lg font-semibold">Internal Libraries</h3>
                    <div className="flex gap-2">
                        <Button
                            type="button"
                            variant="outline"
                            size="sm"
                            onClick={() => window.open("/libraries", "_blank")}
                        >
                            Manage Libraries
                        </Button>
                        <Select
                            value=""
                            onValueChange={(libraryUuid) => {
                                if (libraryUuid === "none") return;
                                const selected = availableLibraries.find((lib) => lib.uuid === libraryUuid);
                                if (selected && !libraries.find((l) => l.uuid === selected.uuid)) {
                                    appendLibrary(selected);
                                }
                            }}
                        >
                            <SelectTrigger className="w-[200px]">
                                <SelectValue placeholder="Add Library" />
                            </SelectTrigger>
                            <SelectContent>
                                {availableLibraries.length === 0 ? (
                                    <SelectItem value="none" disabled>
                                        No libraries found
                                    </SelectItem>
                                ) : availableLibraries.filter(
                                    (lib) => !libraries.find((l) => l.uuid === lib.uuid)
                                ).length === 0 ? (
                                    <SelectItem value="none" disabled>
                                        All libraries added
                                    </SelectItem>
                                ) : (
                                    availableLibraries
                                        .filter((lib) => !libraries.find((l) => l.uuid === lib.uuid))
                                        .map((lib) => (
                                            <SelectItem key={lib.uuid} value={lib.uuid}>
                                                {lib.standardName} ({lib.version})
                                            </SelectItem>
                                        ))
                                )}
                            </SelectContent>
                        </Select>
                    </div>
                </div>

                {libraries.length > 0 && (
                    <div className="grid gap-4">
                        {libraries.map((library, index) => (
                            <Card key={library.id}>
                                <CardContent className="pt-4 flex items-center justify-between">
                                    <div>
                                        <p className="font-medium">{library.standardName}</p>
                                        <p className="text-sm text-muted-foreground">
                                            Version {library.version} • {library.namespace}
                                        </p>
                                    </div>
                                    <Button
                                        type="button"
                                        variant="ghost"
                                        size="sm"
                                        onClick={() => removeLibrary(index)}
                                    >
                                        <Trash2 className="h-4 w-4 text-red-500" />
                                    </Button>
                                </CardContent>
                            </Card>
                        ))}
                    </div>
                )}
            </div>

            {/* Form Actions */}
            <div className="flex gap-2 justify-end">
                {onCancel && (
                    <Button type="button" variant="outline" onClick={onCancel}>
                        Cancel
                    </Button>
                )}
                <Button type="submit" disabled={isLoading}>
                    {isLoading ? "Saving..." : initialData ? "Update Blueprint" : "Create Blueprint"}
                </Button>
            </div>
        </form>
    );
}

interface FunctionalityCardProps {
    funcIndex: number;
    register: any;
    control: any;
    errors: any;
    isExpanded: boolean;
    onToggle: () => void;
    onRemove: () => void;
}

function FunctionalityCard({
    funcIndex,
    register,
    control,
    errors,
    isExpanded,
    onToggle,
    onRemove,
}: FunctionalityCardProps) {
    const {
        fields: operations,
        append: appendOperation,
        remove: removeOperation,
    } = useFieldArray({
        control,
        name: `functionalities.${funcIndex}.operations`,
    });

    // Watch the type field value to update provider dropdown
    const selectedType = useWatch({
        control,
        name: `functionalities.${funcIndex}.type`,
        defaultValue: "",
    });

    const addOperation = () => {
        appendOperation({
            name: "",
            description: "",
        });
    };

    return (
        <Card>
            <CardHeader className="cursor-pointer" onClick={onToggle}>
                <div className="flex items-center justify-between">
                    <CardTitle className="text-base flex items-center gap-2">
                        {isExpanded ? (
                            <ChevronDown className="h-4 w-4" />
                        ) : (
                            <ChevronUp className="h-4 w-4" />
                        )}
                        Functionality {funcIndex + 1}
                    </CardTitle>
                    <Button
                        type="button"
                        variant="ghost"
                        size="sm"
                        onClick={(e) => {
                            e.stopPropagation();
                            onRemove();
                        }}
                    >
                        <Trash2 className="h-4 w-4 text-red-500" />
                    </Button>
                </div>
            </CardHeader>
            {isExpanded && (
                <CardContent className="space-y-4">
                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                            <Label>Category *</Label>
                            <Controller
                                name={`functionalities.${funcIndex}.category`}
                                control={control}
                                render={({ field }) => (
                                    <Select onValueChange={field.onChange} value={field.value}>
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select category" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            {FUNCTIONALITY_CATEGORIES.map((category) => (
                                                <SelectItem key={category} value={category}>
                                                    {category}
                                                </SelectItem>
                                            ))}
                                        </SelectContent>
                                    </Select>
                                )}
                            />
                            {errors?.functionalities?.[funcIndex]?.category && (
                                <p className="text-sm text-red-500">
                                    {errors.functionalities[funcIndex].category.message}
                                </p>
                            )}
                        </div>

                        <div className="space-y-2">
                            <Label>Type *</Label>
                            <Controller
                                name={`functionalities.${funcIndex}.type`}
                                control={control}
                                render={({ field }) => (
                                    <Select
                                        onValueChange={field.onChange}
                                        value={field.value}
                                    >
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select type" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            {FUNCTIONALITY_TYPES.map((type) => (
                                                <SelectItem key={type} value={type}>
                                                    {type}
                                                </SelectItem>
                                            ))}
                                        </SelectContent>
                                    </Select>
                                )}
                            />
                            {errors?.functionalities?.[funcIndex]?.type && (
                                <p className="text-sm text-red-500">
                                    {errors.functionalities[funcIndex].type.message}
                                </p>
                            )}
                        </div>

                        <div className="space-y-2">
                            <Label>Provider</Label>
                            <Controller
                                name={`functionalities.${funcIndex}.provider`}
                                control={control}
                                render={({ field }) => {
                                    const availableProviders = selectedType
                                        ? getProvidersForType(selectedType)
                                        : [];
                                    return (
                                        <Select
                                            onValueChange={field.onChange}
                                            value={field.value}
                                            disabled={!selectedType || availableProviders.length === 0}
                                        >
                                            <SelectTrigger>
                                                <SelectValue
                                                    placeholder={
                                                        selectedType
                                                            ? "Select provider"
                                                            : "Select type first"
                                                    }
                                                />
                                            </SelectTrigger>
                                            <SelectContent>
                                                {availableProviders.map((provider) => (
                                                    <SelectItem key={provider} value={provider}>
                                                        {provider}
                                                    </SelectItem>
                                                ))}
                                            </SelectContent>
                                        </Select>
                                    );
                                }}
                            />
                        </div>

                        <div className="space-y-2">
                            <Label>File Paths (CSV)</Label>
                            <Input
                                placeholder="e.g., /auth/login.ts,/auth/register.ts"
                                {...register(`functionalities.${funcIndex}.filePathsCSV`)}
                            />
                        </div>
                    </div>

                    <div className="flex items-center space-x-2">
                        <Controller
                            name={`functionalities.${funcIndex}.implementsGenerics`}
                            control={control}
                            render={({ field }) => (
                                <Switch
                                    checked={field.value}
                                    onCheckedChange={field.onChange}
                                />
                            )}
                        />
                        <Label>Implements Generics</Label>
                    </div>

                    {/* Operations */}
                    <div className="space-y-2 border-t pt-4">
                        <div className="flex items-center justify-between">
                            <h4 className="font-medium">Operations</h4>
                            <Button
                                type="button"
                                variant="outline"
                                size="sm"
                                onClick={addOperation}
                            >
                                <Plus className="mr-2 h-3 w-3" />
                                Add Operation
                            </Button>
                        </div>

                        {operations.map((operation, opIndex) => (
                            <div
                                key={operation.id}
                                className="flex gap-2 items-start p-3 border rounded-md"
                            >
                                <div className="flex-1 grid grid-cols-2 gap-2">
                                    <div className="space-y-1">
                                        <Label className="text-xs">Name *</Label>
                                        <Input
                                            size="sm"
                                            placeholder="e.g., Login"
                                            {...register(
                                                `functionalities.${funcIndex}.operations.${opIndex}.name`
                                            )}
                                        />
                                        {errors?.functionalities?.[funcIndex]?.operations?.[opIndex]
                                            ?.name && (
                                                <p className="text-xs text-red-500">
                                                    {
                                                        errors.functionalities[funcIndex].operations[
                                                            opIndex
                                                        ].name.message
                                                    }
                                                </p>
                                            )}
                                    </div>
                                    <div className="space-y-1">
                                        <Label className="text-xs">Description</Label>
                                        <Input
                                            size="sm"
                                            placeholder="Describe this operation"
                                            {...register(
                                                `functionalities.${funcIndex}.operations.${opIndex}.description`
                                            )}
                                        />
                                    </div>
                                </div>
                                <Button
                                    type="button"
                                    variant="ghost"
                                    size="sm"
                                    onClick={() => removeOperation(opIndex)}
                                >
                                    <Trash2 className="h-3 w-3 text-red-500" />
                                </Button>
                            </div>
                        ))}
                    </div>
                </CardContent>
            )}
        </Card>
    );
}

