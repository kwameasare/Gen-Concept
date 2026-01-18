import { useState, useCallback, useRef, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import ReactFlow, {
    ReactFlowProvider,
    addEdge,
    useNodesState,
    useEdgesState,
    Controls,
    Background,
    Connection,
    Edge,
    Node,
    MarkerType,
} from "reactflow";
import "reactflow/dist/style.css";
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogFooter,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Save, ArrowLeft, Layers, Zap, Database, Globe, Plus, Trash2, ChevronRight, Home, Eye, Code } from "lucide-react";
import { api } from "@/lib/api";
import { DeleteConfirmDialog } from "@/components/DeleteConfirmDialog";
import { Journey, Operation, JourneyStep, JourneyNode, JourneyEdge } from "@/types/journey";
import { cn } from "@/lib/utils";


interface BreadcrumbItem {
    id: string; // Operation UUID or Step UUID
    label: string;
    steps: JourneyStep[];
}

const initialNodes: Node[] = [];

let id = 0;
const getId = () => `dndnode_${id++}`;

export default function JourneyBuilderPage() {
    const { projectId } = useParams<{ projectId: string }>();
    const navigate = useNavigate();
    const reactFlowWrapper = useRef<HTMLDivElement>(null);
    const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
    const [edges, setEdges, onEdgesChange] = useEdgesState([]);
    const [reactFlowInstance, setReactFlowInstance] = useState<any>(null);
    const [selectedNode, setSelectedNode] = useState<Node | null>(null);

    const [journey, setJourney] = useState<Journey | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [selectedOperationId, setSelectedOperationId] = useState<string | null>(null);
    const [currentParentId, setCurrentParentId] = useState<string | null>(null); // For Zoom/Drill-down
    const [breadcrumbs, setBreadcrumbs] = useState<BreadcrumbItem[]>([]);
    const [viewMode, setViewMode] = useState<"business" | "technical">("business");
    const [deleteJourneyDialog, setDeleteJourneyDialog] = useState(false);

    // Modal State
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [newJourneyName, setNewJourneyName] = useState("");
    const [newEntityId, setNewEntityId] = useState("");
    const [programmingLanguage, setProgrammingLanguage] = useState("Golang");

    // Fetch Graph Data (Nodes & Edges) from Backend
    const fetchGraph = useCallback(async (parentId?: string) => {
        if (!journey || !journey.uuid) return;

        setIsLoading(true);
        setError(null);
        try {
            // Build Query Params
            const params = new URLSearchParams();
            if (parentId) params.append("parent_id", parentId);
            if (viewMode === "business") params.append("level", "HIGH"); // Business view filters for HIGH level nodes

            const res = await api.get<{ result: { nodes: JourneyNode[], edges: JourneyEdge[] } }>(
                `/journeys/${journey.uuid}/graph?${params.toString()}`
            );

            if (res && res.result) {
                loadGraph(res.result.nodes, res.result.edges);
            }
        } catch (error: any) {
            console.error("Failed to fetch graph", error);
            setError("Failed to load journey graph.");
        } finally {
            setIsLoading(false);
        }
    }, [journey, viewMode]);

    // Initial fetch of Journey Metadata
    const fetchJourneyMetadata = useCallback(async () => {
        if (!projectId) return;
        // ... Keep existing logic to find the journey UUID first ...
        try {
            const res = await api.post<{ items: Journey[] }>("/journeys/get-by-filter", {
                pageNumber: 1,
                pageSize: 1,
                dynamicFilter: {
                    field: "project_uuid",
                    operator: "eq",
                    value: projectId,
                },
            });
            if (res.items && res.items.length > 0) {
                setJourney(res.items[0]);
            } else {
                setJourney(null);
            }
        } catch (e) {
            console.error(e);
        }
    }, [projectId]);

    useEffect(() => {
        fetchJourneyMetadata();
    }, [fetchJourneyMetadata]);

    // Fetch graph when journey is loaded or viewMode/parentId changes
    useEffect(() => {
        if (journey) {
            fetchGraph(currentParentId || undefined);
        }
    }, [journey, fetchGraph, currentParentId]);


    const loadGraph = useCallback((jNodes: JourneyNode[], jEdges: JourneyEdge[]) => {
        const newNodes: Node[] = [];
        const newEdges: Edge[] = [];

        // Map JourneyNodes to ReactFlow Nodes
        jNodes.forEach((node) => {
            // Check if node has metadata for position, else default layout
            // For now, simpler auto-layout logic or use stored metadata
            const x = node.metadata?.x || 0;
            const y = node.metadata?.y || 0;

            newNodes.push({
                id: node.uuid,
                type: "default", // or custom node type
                data: {
                    label: node.label,
                    type: node.type,
                    level: node.level,
                    nodeData: node, // Store full object
                },
                position: { x, y }, // We need an auto-layout library if x,y are 0
                style: {
                    background: node.level === "LOW" ? '#f1f5f9' : '#fff',
                    border: '1px solid #e2e8f0',
                    borderRadius: '8px',
                    padding: '10px',
                    width: '200px',
                    textAlign: 'center',
                }
            });
        });

        // Map JourneyEdges to ReactFlow Edges
        jEdges.forEach((edge) => {
            newEdges.push({
                id: edge.uuid,
                source: edge.sourceId,
                target: edge.targetId,
                label: edge.label,
                markerEnd: { type: MarkerType.ArrowClosed },
                type: 'smoothstep'
            });
        });

        // Apply simple layout if positions are missing (Naive vertical stack for MVP)
        // ideally use dagre or elkjs
        if (newNodes.every(n => n.position.x === 0 && n.position.y === 0)) {
            newNodes.forEach((node, index) => {
                node.position = { x: 250, y: index * 100 };
            });
        }

        setNodes(newNodes);
        setEdges(newEdges);
    }, [setNodes, setEdges]);


    const handleOperationSelect = (_entityId: string, op: Operation) => {
        setSelectedOperationId(op.uuid);
        setSelectedNode(null);
        setBreadcrumbs([{
            id: op.uuid,
            label: op.name,
            steps: []
        }]);
        // When selecting an operation, we treat it as zooming into that operation node (or root)
        // Ideally, the operation UUID is the parent ID in the graph
        setCurrentParentId(op.uuid);
    };

    // Reload flow when viewMode changes - handled by fetchGraph dependency on viewMode
    // useEffect(() => {
    //    if (selectedOperationId) {
    //        const currentSteps = breadcrumbs.length > 0 ? breadcrumbs[breadcrumbs.length - 1].steps : [];
    //        loadFlow(currentSteps, viewMode);
    //    }
    // }, [viewMode, loadFlow, selectedOperationId]); 


    // Double Click to Drill Down
    const handleNodeDoubleClick = useCallback((_event: React.MouseEvent, node: Node) => {
        const jNode = node.data.nodeData as JourneyNode;
        // Assume double click means "Drill Down" if it's a composite node or based on type
        // Or simply set parentId to this node's UUID to find children

        // Update breadcrumbs
        const newCrumb: BreadcrumbItem = {
            id: jNode.uuid,
            label: jNode.label || jNode.type,
            steps: [] // We don't need steps anymore, graph fetch handles it
        };
        setBreadcrumbs([...breadcrumbs, newCrumb]);

        setCurrentParentId(jNode.uuid);
        setSelectedNode(null);

    }, [breadcrumbs]);

    const handleBreadcrumbClick = (index: number) => {
        // Go back to that level
        if (index === -1) {
            // Home
            setCurrentParentId(null);
            setBreadcrumbs([]);
        } else {
            const targetCrumb = breadcrumbs[index];
            setCurrentParentId(targetCrumb.id);
            setBreadcrumbs(breadcrumbs.slice(0, index + 1));
        }
        setSelectedNode(null);
    };

    const onConnect = useCallback(
        (params: Connection | Edge) => setEdges((eds) => addEdge(params, eds)),
        [setEdges]
    );

    const onNodeClick = useCallback((_: React.MouseEvent, node: Node) => {
        setSelectedNode(node);
    }, []);

    const onDragOver = useCallback((event: React.DragEvent) => {
        event.preventDefault();
        event.dataTransfer.dropEffect = "move";
    }, []);

    const onDrop = useCallback(
        (event: React.DragEvent) => {
            event.preventDefault();

            const type = event.dataTransfer.getData("application/reactflow");

            if (typeof type === "undefined" || !type) {
                return;
            }

            const position = reactFlowInstance.project({
                x: event.clientX - (reactFlowWrapper.current?.getBoundingClientRect().left || 0),
                y: event.clientY - (reactFlowWrapper.current?.getBoundingClientRect().top || 0),
            });
            const newNode = {
                id: getId(),
                type: "default",
                position,
                data: { label: `${type} node` },
            };

            setNodes((nds) => nds.concat(newNode));
        },
        [reactFlowInstance, setNodes]
    );

    const handleInitializeJourney = async () => {
        if (!projectId) return;

        setIsLoading(true);
        setError(null);
        try {
            // Default values for new journey
            const payload = {
                projectUuid: projectId,
                programmingLanguage: "Golang", // Default
                entityJourneys: []
            };

            const res = await api.post<{ result: Journey }>("/journeys/", payload);
            if (res) {
                // Refresh journey data
                fetchJourneyMetadata();
            }
        } catch (error: any) {
            console.error("Failed to initialize journey", error);
            setError("Failed to create journey. Please try again.");
            setIsLoading(false); // Stop loading if error, otherwise fetchJourney will handle it
        }
    };

    const handleOpenModal = () => {
        setNewJourneyName("");
        setNewEntityId("");
        setProgrammingLanguage(journey?.programmingLanguage || "Golang");
        setIsModalOpen(true);
    };

    const handleCreateJourney = async () => {
        if (!projectId || !newJourneyName || !newEntityId) return;

        let currentJourney = journey;

        // 1. If no journey exists, create one first
        if (!currentJourney) {
            try {
                setIsLoading(true);
                const payload = {
                    projectUuid: projectId,
                    programmingLanguage: "Golang",
                    entityJourneys: []
                };
                const res = await api.post<{ result: Journey }>("/journeys/", payload);
                if (res && res.result) {
                    currentJourney = res.result;
                    setJourney(res.result);
                } else {
                    throw new Error("Failed to initialize journey container");
                }
            } catch (error) {
                console.error("Failed to initialize journey", error);
                setError("Failed to initialize journey. Please try again.");
                setIsLoading(false);
                return;
            }
        }

        // 2. Create a new EntityJourney (Logical Journey)
        const newEntityName = newJourneyName;
        // We use the provided Entity ID
        const entityId = newEntityId;
        const newOperationName = "Flow 1";

        let updatedJourney = { ...currentJourney };

        // Update programming language
        updatedJourney.programmingLanguage = programmingLanguage;

        if (!updatedJourney.entityJourneys) {
            updatedJourney.entityJourneys = [];
        }

        // Create new EntityJourney
        const newEntity = {
            uuid: crypto.randomUUID(),
            entityName: newEntityName,
            entityId: entityId,
            operations: []
        } as any;

        // Create initial operation for this journey
        const newOperation = {
            uuid: crypto.randomUUID(),
            name: newOperationName,
            type: "CUSTOM_API",
            description: "",
            frontendJourney: [],
            backendJourney: []
        };

        newEntity.operations.push(newOperation);
        updatedJourney.entityJourneys.push(newEntity);

        try {
            setIsLoading(true);
            const res = await api.put<{ result: Journey }>(`/journeys/${currentJourney.uuid}`, updatedJourney);

            if (res) {
                await fetchJourneyMetadata();
                // Auto-select the new operation
                handleOperationSelect(newEntity.uuid, newOperation);
            }
        } catch (error) {
            console.error("Failed to add journey", error);
            setError("Failed to save new journey.");
        } finally {
            setIsLoading(false);
            setIsModalOpen(false);
        }
    };

    const handleDeleteJourney = async () => {
        if (!journey) return;
        try {
            await api.delete(`/journeys/${journey.uuid}`);
            navigate(`/projects/${projectId}`);
        } catch (error) {
            console.error("Failed to delete journey", error);
            setError("Failed to delete journey.");
        }
    };

    return (
        <div className="h-screen flex flex-col">
            <header className="h-14 border-b flex items-center justify-between px-6 bg-card z-10">
                <div className="flex items-center gap-4">
                    <Button variant="ghost" size="sm" onClick={() => navigate(`/projects/${projectId}`)}>
                        <ArrowLeft className="w-4 h-4 mr-2" /> Back
                    </Button>
                    <h1 className="font-semibold">Journey Builder</h1>
                </div>
                <div className="flex gap-2">
                    <Button size="sm">
                        <Save className="w-4 h-4 mr-2" /> Save Journey
                    </Button>
                    <div className="border-l pl-2 ml-2 flex items-center">
                        <div className="bg-muted p-1 rounded-md flex gap-1">
                            <button
                                onClick={() => setViewMode("business")}
                                className={cn(
                                    "p-1.5 rounded-sm text-xs font-medium flex items-center gap-1 transition-all",
                                    viewMode === "business" ? "bg-white shadow-sm text-primary" : "text-muted-foreground hover:text-foreground"
                                )}
                                title="Business View (High Level)"
                            >
                                <Eye className="w-3 h-3" /> Business
                            </button>
                            <button
                                onClick={() => setViewMode("technical")}
                                className={cn(
                                    "p-1.5 rounded-sm text-xs font-medium flex items-center gap-1 transition-all",
                                    viewMode === "technical" ? "bg-white shadow-sm text-primary" : "text-muted-foreground hover:text-foreground"
                                )}
                                title="Technical View (All Details)"
                            >
                                <Code className="w-3 h-3" /> Technical
                            </button>
                        </div>
                    </div>
                    {journey && (
                        <Button
                            variant="destructive"
                            size="sm"
                            onClick={() => setDeleteJourneyDialog(true)}
                        >
                            <Trash2 className="w-4 h-4 mr-2" /> Delete Journey
                        </Button>
                    )}
                </div>
            </header>
            <div className="flex-1 flex overflow-hidden">
                <aside className="w-72 border-r bg-muted flex flex-col z-10 overflow-y-auto shrink-0">
                    <div className="p-4 border-b bg-card flex items-center justify-between">
                        <div>
                            <h2 className="font-semibold mb-1">Explorer</h2>
                            <p className="text-xs text-muted-foreground">Select an operation to edit</p>
                        </div>
                        <Button variant="ghost" size="icon" onClick={handleOpenModal}>
                            <Plus className="w-4 h-4" />
                        </Button>
                    </div>
                    <div className="p-2 space-y-2 flex-1">
                        {isLoading ? (
                            <div className="p-4 text-center text-sm text-muted-foreground">
                                Loading journey...
                            </div>
                        ) : error ? (
                            <div className="p-4 text-center text-sm text-red-500">
                                {error}
                                <Button variant="outline" size="sm" className="mt-2" onClick={fetchJourneyMetadata}>Retry</Button>
                            </div>
                        ) : !journey ? (
                            <div className="p-4 text-center text-sm text-muted-foreground">
                                No journey configuration found.
                                <Button
                                    variant="outline"
                                    size="sm"
                                    className="mt-2"
                                    onClick={handleInitializeJourney}
                                >
                                    Initialize Journey
                                </Button>
                            </div>
                        ) : (
                            journey.entityJourneys?.map((ej) => (
                                <div key={ej.uuid} className="space-y-1">
                                    <div className="px-2 py-1.5 text-sm font-medium flex items-center gap-2 text-gray-700 bg-gray-100 rounded-md mb-1">
                                        <Layers className="w-4 h-4" />
                                        {ej.entityName}
                                    </div>
                                    <div className="pl-4 space-y-1">
                                        {ej.operations?.map((op) => (
                                            <button
                                                key={op.uuid}
                                                onClick={() => handleOperationSelect(ej.entityId, op)}
                                                className={cn(
                                                    "w-full text-left px-2 py-1.5 text-sm rounded-md flex items-center gap-2 transition-colors",
                                                    selectedOperationId === op.uuid
                                                        ? "bg-blue-100 text-blue-700 font-medium"
                                                        : "hover:bg-gray-200 text-gray-600"
                                                )}
                                            >
                                                <Zap className="w-3 h-3" />
                                                {op.name}
                                            </button>
                                        ))}
                                    </div>
                                </div>
                            ))
                        )}
                    </div>

                    <div className={cn("mt-auto p-4 border-t bg-card", !selectedOperationId && "opacity-50 pointer-events-none")}>
                        <div className="text-sm font-medium mb-2">Toolbox</div>
                        <p className="text-[10px] text-muted-foreground mb-2">Drag items to the canvas</p>
                        <div className="grid grid-cols-2 gap-2">
                            <div
                                className="bg-card p-2 rounded border shadow-sm cursor-move text-xs flex flex-col items-center gap-1 hover:border-blue-500 transition-colors"
                                onDragStart={(event) => event.dataTransfer.setData("application/reactflow", "API_CALL")}
                                draggable
                            >
                                <Globe className="w-4 h-4 text-blue-500" />
                                API Call
                            </div>
                            <div
                                className="bg-card p-2 rounded border shadow-sm cursor-move text-xs flex flex-col items-center gap-1 hover:border-blue-500 transition-colors"
                                onDragStart={(event) => event.dataTransfer.setData("application/reactflow", "DATABASE_OPERATION")}
                                draggable
                            >
                                <Database className="w-4 h-4 text-purple-500" />
                                DB Op
                            </div>
                        </div>
                    </div>
                </aside>

                <div className="flex-1 h-full relative" ref={reactFlowWrapper}>
                    {selectedOperationId ? (
                        <ReactFlowProvider>
                            <ReactFlow
                                nodes={nodes}
                                edges={edges}
                                onNodesChange={onNodesChange}
                                onEdgesChange={onEdgesChange}
                                onConnect={onConnect}
                                onNodeClick={onNodeClick}
                                onNodeDoubleClick={handleNodeDoubleClick}
                                onInit={setReactFlowInstance}
                                onDrop={onDrop}
                                onDragOver={onDragOver}
                                fitView
                            >
                                <Controls />
                                <Background />
                                <div className="absolute top-4 left-4 z-10 bg-white/90 p-2 rounded-md shadow-sm border flex items-center gap-2 text-sm">
                                    {/* Home Button */}
                                    <div className="flex items-center">
                                        <button
                                            onClick={() => handleBreadcrumbClick(-1)}
                                            className={cn("hover:underline flex items-center gap-1", breadcrumbs.length === 0 ? "font-bold text-primary" : "text-muted-foreground")}
                                        >
                                            <Home className="w-3 h-3" /> Home
                                        </button>
                                    </div>

                                    {breadcrumbs.map((crumb, index) => (
                                        <div key={crumb.id} className="flex items-center">
                                            <ChevronRight className="w-4 h-4 text-muted-foreground mx-1" />
                                            <button
                                                onClick={() => handleBreadcrumbClick(index)}
                                                className={cn(
                                                    "hover:underline flex items-center gap-1",
                                                    index === breadcrumbs.length - 1 ? "font-bold text-primary" : "text-muted-foreground"
                                                )}
                                            >
                                                {crumb.label}
                                            </button>
                                        </div>
                                    ))}
                                </div>
                            </ReactFlow>
                        </ReactFlowProvider>
                    ) : (
                        <div className="absolute inset-0 flex items-center justify-center text-muted-foreground bg-muted/50">
                            <div className="text-center">
                                <Layers className="w-12 h-12 mx-auto mb-4 opacity-20" />
                                <p>Select an operation from the sidebar to view its flow</p>
                            </div>
                        </div>
                    )}
                </div>

                {selectedNode && selectedNode.data.stepData && (
                    <aside className="w-80 border-l bg-card p-4 overflow-y-auto shadow-xl z-20">
                        <div className="flex items-center justify-between mb-4">
                            <h3 className="font-semibold text-lg">Step Details</h3>
                            <Button variant="ghost" size="sm" onClick={() => setSelectedNode(null)}>
                                Close
                            </Button>
                        </div>

                        <div className="space-y-4">
                            <div>
                                <label className="text-xs font-medium text-muted-foreground">Type</label>
                                <div className="text-sm font-medium">{selectedNode.data.stepData.type}</div>
                            </div>

                            {selectedNode.data.stepData.description && (
                                <div>
                                    <label className="text-xs font-medium text-muted-foreground">Description</label>
                                    <div className="text-sm text-gray-600">{selectedNode.data.stepData.description}</div>
                                </div>
                            )}

                            {selectedNode.data.stepData.condition && (
                                <div>
                                    <label className="text-xs font-medium text-muted-foreground">Condition</label>
                                    <div className="text-sm font-mono bg-muted p-2 rounded border mt-1">
                                        {selectedNode.data.stepData.condition}
                                    </div>
                                </div>
                            )}

                            {selectedNode.data.stepData.curl && (
                                <div>
                                    <label className="text-xs font-medium text-muted-foreground">cURL</label>
                                    <div className="text-xs font-mono bg-muted p-2 rounded border mt-1 break-all">
                                        {selectedNode.data.stepData.curl}
                                    </div>
                                </div>
                            )}

                            {selectedNode.data.stepData.fieldsInvolved?.length > 0 && (
                                <div>
                                    <label className="text-xs font-medium text-muted-foreground">Fields Involved</label>
                                    <div className="flex flex-wrap gap-1 mt-1">
                                        {selectedNode.data.stepData.fieldsInvolved.map((f: any) => (
                                            <span key={f.id} className="text-xs bg-blue-50 text-blue-700 px-2 py-0.5 rounded border border-blue-100">
                                                {f.name}
                                            </span>
                                        ))}
                                    </div>
                                </div>
                            )}

                            {selectedNode.data.stepData.responseActions?.length > 0 && (
                                <div>
                                    <label className="text-xs font-medium text-muted-foreground">Response Actions</label>
                                    <div className="space-y-2 mt-1">
                                        {selectedNode.data.stepData.responseActions.map((action: any, idx: number) => (
                                            <div key={idx} className="text-xs bg-muted p-2 rounded border">
                                                <div className="font-medium text-purple-700">{action.type}</div>
                                                {action.fieldId && <div>Field: {action.fieldId}</div>}
                                                {action.value && <div>Value: {action.value}</div>}
                                            </div>
                                        ))}
                                    </div>
                                </div>
                            )}

                            <div className="pt-4 border-t">
                                <pre className="text-[10px] text-gray-400 overflow-x-auto">
                                    {JSON.stringify(selectedNode.data.stepData, null, 2)}
                                </pre>
                            </div>

                            {selectedNode.data.stepData.subSteps && selectedNode.data.stepData.subSteps.length > 0 && (
                                <div className="bg-blue-50 border border-blue-100 p-2 rounded text-xs text-blue-700">
                                    Contains {selectedNode.data.stepData.subSteps.length} sub-steps. Double-click node to view.
                                </div>
                            )}
                        </div>
                    </aside>
                )}
            </div>
            {/* Create Journey Modal */}
            <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Create New Journey</DialogTitle>
                    </DialogHeader>
                    <div className="space-y-4 py-4">
                        <div className="space-y-2">
                            <Label htmlFor="journeyName">Journey Name</Label>
                            <Input
                                id="journeyName"
                                placeholder="e.g., User Onboarding"
                                value={newJourneyName}
                                onChange={(e) => setNewJourneyName(e.target.value)}
                            />
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="entityId">Entity ID</Label>
                            <Input
                                id="entityId"
                                placeholder="e.g., user_onboarding_flow"
                                value={newEntityId}
                                onChange={(e) => setNewEntityId(e.target.value)}
                            />
                            <p className="text-xs text-muted-foreground">
                                A unique identifier for this journey flow.
                            </p>
                        </div>
                        <div className="space-y-2">
                            <Label htmlFor="programmingLanguage">Programming Language</Label>
                            <Select value={programmingLanguage} onValueChange={setProgrammingLanguage}>
                                <SelectTrigger>
                                    <SelectValue placeholder="Select language" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem value="Golang">Golang</SelectItem>
                                    <SelectItem value="Python">Python</SelectItem>
                                    <SelectItem value="Java">Java</SelectItem>
                                    <SelectItem value="TypeScript">TypeScript</SelectItem>
                                    <SelectItem value="JavaScript">JavaScript</SelectItem>
                                    <SelectItem value="Csharp">C#</SelectItem>
                                    <SelectItem value="Rust">Rust</SelectItem>
                                    <SelectItem value="Php">PHP</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                    </div>
                    <DialogFooter>
                        <Button variant="outline" onClick={() => setIsModalOpen(false)}>
                            Cancel
                        </Button>
                        <Button onClick={handleCreateJourney} disabled={!newJourneyName || !newEntityId || isLoading}>
                            {isLoading ? "Creating..." : "Create Journey"}
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>

            {/* Delete Confirmation Dialog */}
            <DeleteConfirmDialog
                open={deleteJourneyDialog}
                onOpenChange={setDeleteJourneyDialog}
                onConfirm={handleDeleteJourney}
                title="Delete Journey"
                description="Are you sure you want to delete this entire journey configuration? This action cannot be undone."
                itemName={journey?.programmingLanguage ? `${journey.programmingLanguage} Journey` : "Journey"}
            />
        </div >
    );
}
