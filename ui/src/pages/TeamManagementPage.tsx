import { useEffect, useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { api, getUser } from "@/lib/api";
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
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import type { Team, CreateTeamRequest } from "@/types/team";

export default function TeamManagementPage() {
    const [teams, setTeams] = useState<Team[]>([]);
    const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    // Get current org from user context
    const user = getUser();
    // Go struct fields are usually PascalCase in JSON if no tags, but checking both just in case
    const currentOrgId = user?.OrganizationID || user?.organizationID || 1;

    const [newTeam, setNewTeam] = useState<Partial<CreateTeamRequest>>({
        name: "",
        description: "",
        organizationID: currentOrgId,
    });

    useEffect(() => {
        fetchTeams();
    }, []);

    const fetchTeams = async () => {
        try {
            const response: any = await api.post("/teams/get-by-filter", {
                pageNumber: 1,
                pageSize: 100,
                // Add filter by OrgID here if backend supports it in dynamic filter
            });
            setTeams(response.items || []);
        } catch (error) {
            console.error("Error fetching teams:", error);
        }
    };

    const handleCreateTeam = async () => {
        if (!newTeam.name) return;

        setIsLoading(true);
        try {
            await api.post("/teams", {
                ...newTeam,
                organizationID: currentOrgId // Ensure Org ID is set
            });
            setIsCreateDialogOpen(false);
            setNewTeam({
                name: "",
                description: "",
                organizationID: currentOrgId,
            });
            fetchTeams();
        } catch (error) {
            console.error("Error creating team:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const handleDeleteTeam = async (uuid: string) => {
        try {
            await api.delete(`/teams/${uuid}`);
            fetchTeams();
        } catch (error) {
            console.error("Error deleting team:", error);
        }
    };

    return (
        <div className="min-h-screen bg-background p-8">
            <div className="max-w-6xl mx-auto space-y-8">
                <div className="flex items-center justify-between">
                    <div>
                        <h1 className="text-3xl font-bold tracking-tight">Team Management</h1>
                        <p className="text-muted-foreground">
                            Manage your organization's teams
                        </p>
                    </div>

                    <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
                        <DialogTrigger asChild>
                            <Button>
                                <Plus className="mr-2 h-4 w-4" />
                                Create Team
                            </Button>
                        </DialogTrigger>
                        <DialogContent>
                            <DialogHeader>
                                <DialogTitle>Create New Team</DialogTitle>
                                <DialogDescription>
                                    Add a new team to your organization
                                </DialogDescription>
                            </DialogHeader>

                            <div className="space-y-4">
                                <div className="space-y-2">
                                    <Label>Team Name *</Label>
                                    <Input
                                        placeholder="e.g., Frontend Team"
                                        value={newTeam.name}
                                        onChange={(e) =>
                                            setNewTeam({ ...newTeam, name: e.target.value })
                                        }
                                    />
                                </div>

                                <div className="space-y-2">
                                    <Label>Description</Label>
                                    <Textarea
                                        placeholder="What does this team do?"
                                        value={newTeam.description}
                                        onChange={(e) =>
                                            setNewTeam({ ...newTeam, description: e.target.value })
                                        }
                                    />
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
                                <Button onClick={handleCreateTeam} disabled={isLoading}>
                                    Create Team
                                </Button>
                            </DialogFooter>
                        </DialogContent>
                    </Dialog>
                </div>

                <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                    {teams.map((team) => (
                        <Card key={team.uuid}>
                            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                                <CardTitle className="text-xl font-bold">{team.name}</CardTitle>
                                <Button
                                    variant="ghost"
                                    size="sm"
                                    onClick={() => handleDeleteTeam(team.uuid)}
                                >
                                    <Trash2 className="h-4 w-4 text-red-500" />
                                </Button>
                            </CardHeader>
                            <CardContent>
                                <p className="text-sm text-muted-foreground">{team.description}</p>
                            </CardContent>
                        </Card>
                    ))}

                    {teams.length === 0 && (
                        <Card className="col-span-full">
                            <CardContent className="py-12 text-center text-muted-foreground">
                                No teams yet. Create your first team to get started.
                            </CardContent>
                        </Card>
                    )}
                </div>
            </div>
        </div>
    );
}
