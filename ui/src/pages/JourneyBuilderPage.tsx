import { useState, useCallback, useRef } from "react";
import ReactFlow, {
    ReactFlowProvider,
    addEdge,
    useNodesState,
    useEdgesState,
    Controls,
    Background,
    Connection,
    Edge,
} from "reactflow";
import "reactflow/dist/style.css";
import { Button } from "@/components/ui/button";
import { Save } from "lucide-react";

const initialNodes = [
    {
        id: "1",
        type: "input",
        data: { label: "Start Journey" },
        position: { x: 250, y: 5 },
    },
];

let id = 0;
const getId = () => `dndnode_${id++}`;

export default function JourneyBuilderPage() {
    const reactFlowWrapper = useRef<HTMLDivElement>(null);
    const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
    const [edges, setEdges, onEdgesChange] = useEdgesState([]);
    const [reactFlowInstance, setReactFlowInstance] = useState<any>(null);

    const onConnect = useCallback(
        (params: Connection | Edge) => setEdges((eds) => addEdge(params, eds)),
        [setEdges]
    );

    const onDragOver = useCallback((event: React.DragEvent) => {
        event.preventDefault();
        event.dataTransfer.dropEffect = "move";
    }, []);

    const onDrop = useCallback(
        (event: React.DragEvent) => {
            event.preventDefault();

            const type = event.dataTransfer.getData("application/reactflow");

            // check if the dropped element is valid
            if (typeof type === "undefined" || !type) {
                return;
            }

            const position = reactFlowInstance.project({
                x: event.clientX - (reactFlowWrapper.current?.getBoundingClientRect().left || 0),
                y: event.clientY - (reactFlowWrapper.current?.getBoundingClientRect().top || 0),
            });
            const newNode = {
                id: getId(),
                type,
                position,
                data: { label: `${type} node` },
            };

            setNodes((nds) => nds.concat(newNode));
        },
        [reactFlowInstance, setNodes]
    );

    return (
        <div className="h-screen flex flex-col">
            <header className="h-14 border-b flex items-center justify-between px-6 bg-white z-10">
                <h1 className="font-semibold">Journey Builder</h1>
                <Button size="sm">
                    <Save className="w-4 h-4 mr-2" /> Save Journey
                </Button>
            </header>
            <div className="flex-1 flex overflow-hidden">
                <aside className="w-64 border-r bg-gray-50 p-4 flex flex-col gap-4 z-10">
                    <div className="text-sm font-medium text-muted-foreground">
                        Operations
                    </div>
                    <div
                        className="bg-white p-3 rounded border shadow-sm cursor-move flex items-center gap-2"
                        onDragStart={(event) =>
                            event.dataTransfer.setData("application/reactflow", "default")
                        }
                        draggable
                    >
                        <div className="w-3 h-3 rounded-full bg-blue-500" />
                        Generic Step
                    </div>
                    <div
                        className="bg-white p-3 rounded border shadow-sm cursor-move flex items-center gap-2"
                        onDragStart={(event) =>
                            event.dataTransfer.setData("application/reactflow", "output")
                        }
                        draggable
                    >
                        <div className="w-3 h-3 rounded-full bg-green-500" />
                        End Journey
                    </div>
                </aside>
                <div className="flex-1 h-full" ref={reactFlowWrapper}>
                    <ReactFlowProvider>
                        <ReactFlow
                            nodes={nodes}
                            edges={edges}
                            onNodesChange={onNodesChange}
                            onEdgesChange={onEdgesChange}
                            onConnect={onConnect}
                            onInit={setReactFlowInstance}
                            onDrop={onDrop}
                            onDragOver={onDragOver}
                            fitView
                        >
                            <Controls />
                            <Background />
                        </ReactFlow>
                    </ReactFlowProvider>
                </div>
            </div>
        </div>
    );
}
