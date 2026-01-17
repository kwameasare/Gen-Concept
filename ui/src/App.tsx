import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import OnboardingPage from "./pages/OnboardingPage";
import DashboardPage from "./pages/DashboardPage";
import ProjectDetailPage from "./pages/ProjectDetailPage";
import EntityDetailPage from "./pages/EntityDetailPage";
import JourneyBuilderPage from "./pages/JourneyBuilderPage";
import BlueprintDetailPage from "./pages/BlueprintDetailPage";
import LibraryManagementPage from "./pages/LibraryManagementPage";
import TeamManagementPage from "./pages/TeamManagementPage";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<LoginPage />} />
                <Route path="/onboard" element={<OnboardingPage />} />
                <Route path="/dashboard" element={<DashboardPage />} />
                <Route path="/projects/:id" element={<ProjectDetailPage />} />
                <Route path="/projects/:projectId/entities/:entityId" element={<EntityDetailPage />} />
                <Route path="/projects/:projectId/journeys/builder" element={<JourneyBuilderPage />} />
                <Route path="/blueprints/:id" element={<BlueprintDetailPage />} />
                <Route path="/blueprints/:id" element={<BlueprintDetailPage />} />
                <Route path="/libraries" element={<LibraryManagementPage />} />
                <Route path="/teams" element={<TeamManagementPage />} />
                <Route path="/" element={<Navigate to="/login" replace />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
