import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useNavigate, Link } from "react-router-dom";
import { motion } from "framer-motion";
import { api, setToken, setUser } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Building2, Sparkles, Zap, Rocket } from "lucide-react";

const schema = z.object({
    organization: z.object({
        name: z.string().min(2, "Organization name is required"),
        domain: z.string().min(3, "Domain is required"),
        description: z.string().optional(),
    }),
    user: z.object({
        username: z.string().min(3, "Username is required"),
        email: z.string().email("Invalid email address"),
        password: z.string().min(6, "Password must be at least 6 characters"),
    }),
});

type FormData = z.infer<typeof schema>;

export default function OnboardingPage() {
    const navigate = useNavigate();
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<FormData>({
        resolver: zodResolver(schema),
    });

    const onSubmit = async (data: FormData) => {
        setLoading(true);
        setError(null);
        try {
            await api.post("/organizations/onboard", data);

            // Login to get token
            const tokenData = await api.post<{ accessToken: string }>("/users/login-by-username", {
                username: data.user.username,
                password: data.user.password,
            });

            setToken(tokenData.accessToken);
            if ((tokenData as any).user) {
                setUser((tokenData as any).user);
            }
            navigate("/dashboard");
        } catch (err: any) {
            setError(err.message || "Failed to onboard");
        } finally {
            setLoading(false);
        }
    };

    const features = [
        { icon: Zap, title: "Lightning Fast", desc: "Build projects in minutes" },
        { icon: Rocket, title: "Production Ready", desc: "Enterprise-grade solutions" },
        { icon: Sparkles, title: "AI Powered", desc: "Smart code generation" },
    ];

    return (
        <div className="min-h-screen flex overflow-hidden">
            {/* Left Side - Form */}
            <div className="flex-1 flex items-center justify-center p-8 relative overflow-hidden">
                {/* Animated Background Gradient */}
                <div className="absolute inset-0 bg-gradient-to-br from-violet-500/10 via-fuchsia-500/10 to-pink-500/10">
                    <div className="absolute inset-0 bg-[radial-gradient(circle_at_50%_120%,rgba(120,119,198,0.3),rgba(255,255,255,0))]" />
                </div>

                {/* Floating Orbs */}
                <motion.div
                    className="absolute top-20 left-20 w-72 h-72 bg-violet-500/30 rounded-full blur-3xl"
                    animate={{
                        scale: [1, 1.2, 1],
                        opacity: [0.3, 0.5, 0.3],
                    }}
                    transition={{
                        duration: 8,
                        repeat: Infinity,
                        ease: "easeInOut",
                    }}
                />
                <motion.div
                    className="absolute bottom-20 right-20 w-96 h-96 bg-fuchsia-500/30 rounded-full blur-3xl"
                    animate={{
                        scale: [1, 1.3, 1],
                        opacity: [0.3, 0.4, 0.3],
                    }}
                    transition={{
                        duration: 10,
                        repeat: Infinity,
                        ease: "easeInOut",
                        delay: 1,
                    }}
                />

                {/* Form Container */}
                <motion.div
                    initial={{ opacity: 0, scale: 0.95 }}
                    animate={{ opacity: 1, scale: 1 }}
                    transition={{ duration: 0.5 }}
                    className="w-full max-w-md relative z-10"
                >
                    {/* Glassmorphic Card */}
                    <div className="backdrop-blur-xl bg-white/70 dark:bg-gray-900/70 rounded-3xl shadow-2xl border border-white/20 p-8">
                        {/* Header */}
                        <motion.div
                            initial={{ opacity: 0, y: -20 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.2 }}
                            className="text-center mb-8"
                        >
                            <div className="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-gradient-to-br from-violet-500 to-fuchsia-500 mb-4">
                                <Building2 className="w-8 h-8 text-white" />
                            </div>
                            <h1 className="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-violet-600 to-fuchsia-600 mb-2">
                                Welcome to Gen-Concept
                            </h1>
                            <p className="text-gray-600 dark:text-gray-300">
                                Create your organization and start building
                            </p>
                        </motion.div>

                        {/* Form */}
                        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
                            {error && (
                                <motion.div
                                    initial={{ opacity: 0, x: -20 }}
                                    animate={{ opacity: 1, x: 0 }}
                                    className="p-4 text-sm text-red-600 bg-red-50 dark:bg-red-900/20 rounded-xl border border-red-200 dark:border-red-800"
                                >
                                    {error}
                                </motion.div>
                            )}

                            {/* Organization Section */}
                            <motion.div
                                initial={{ opacity: 0, y: 20 }}
                                animate={{ opacity: 1, y: 0 }}
                                transition={{ delay: 0.3 }}
                                className="space-y-4"
                            >
                                <div className="flex items-center gap-2 mb-3">
                                    <div className="w-1 h-5 bg-gradient-to-b from-violet-500 to-fuchsia-500 rounded-full" />
                                    <h3 className="text-sm font-semibold text-gray-700 dark:text-gray-200 uppercase tracking-wider">
                                        Organization Details
                                    </h3>
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="orgName" className="text-gray-700 dark:text-gray-200">
                                        Organization Name
                                    </Label>
                                    <Input
                                        id="orgName"
                                        placeholder="Acme Corporation"
                                        {...register("organization.name")}
                                        className="border-gray-200 dark:border-gray-700 focus:border-violet-500 dark:focus:border-violet-400 transition-all duration-300 hover:border-violet-300"
                                    />
                                    {errors.organization?.name && (
                                        <motion.p
                                            initial={{ opacity: 0, y: -10 }}
                                            animate={{ opacity: 1, y: 0 }}
                                            className="text-sm text-red-500"
                                        >
                                            {errors.organization.name.message}
                                        </motion.p>
                                    )}
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="domain" className="text-gray-700 dark:text-gray-200">
                                        Domain
                                    </Label>
                                    <Input
                                        id="domain"
                                        placeholder="acme.com"
                                        {...register("organization.domain")}
                                        className="border-gray-200 dark:border-gray-700 focus:border-violet-500 dark:focus:border-violet-400 transition-all duration-300 hover:border-violet-300"
                                    />
                                    {errors.organization?.domain && (
                                        <motion.p
                                            initial={{ opacity: 0, y: -10 }}
                                            animate={{ opacity: 1, y: 0 }}
                                            className="text-sm text-red-500"
                                        >
                                            {errors.organization.domain.message}
                                        </motion.p>
                                    )}
                                </div>
                            </motion.div>

                            {/* Admin Account Section */}
                            <motion.div
                                initial={{ opacity: 0, y: 20 }}
                                animate={{ opacity: 1, y: 0 }}
                                transition={{ delay: 0.4 }}
                                className="space-y-4"
                            >
                                <div className="flex items-center gap-2 mb-3">
                                    <div className="w-1 h-5 bg-gradient-to-b from-violet-500 to-fuchsia-500 rounded-full" />
                                    <h3 className="text-sm font-semibold text-gray-700 dark:text-gray-200 uppercase tracking-wider">
                                        Admin Account
                                    </h3>
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="username" className="text-gray-700 dark:text-gray-200">
                                        Username
                                    </Label>
                                    <Input
                                        id="username"
                                        placeholder="admin"
                                        {...register("user.username")}
                                        className="border-gray-200 dark:border-gray-700 focus:border-violet-500 dark:focus:border-violet-400 transition-all duration-300 hover:border-violet-300"
                                    />
                                    {errors.user?.username && (
                                        <motion.p
                                            initial={{ opacity: 0, y: -10 }}
                                            animate={{ opacity: 1, y: 0 }}
                                            className="text-sm text-red-500"
                                        >
                                            {errors.user.username.message}
                                        </motion.p>
                                    )}
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="email" className="text-gray-700 dark:text-gray-200">
                                        Email
                                    </Label>
                                    <Input
                                        id="email"
                                        type="email"
                                        placeholder="admin@acme.com"
                                        {...register("user.email")}
                                        className="border-gray-200 dark:border-gray-700 focus:border-violet-500 dark:focus:border-violet-400 transition-all duration-300 hover:border-violet-300"
                                    />
                                    {errors.user?.email && (
                                        <motion.p
                                            initial={{ opacity: 0, y: -10 }}
                                            animate={{ opacity: 1, y: 0 }}
                                            className="text-sm text-red-500"
                                        >
                                            {errors.user.email.message}
                                        </motion.p>
                                    )}
                                </div>

                                <div className="space-y-2">
                                    <Label htmlFor="password" className="text-gray-700 dark:text-gray-200">
                                        Password
                                    </Label>
                                    <Input
                                        id="password"
                                        type="password"
                                        placeholder="••••••••"
                                        {...register("user.password")}
                                        className="border-gray-200 dark:border-gray-700 focus:border-violet-500 dark:focus:border-violet-400 transition-all duration-300 hover:border-violet-300"
                                    />
                                    {errors.user?.password && (
                                        <motion.p
                                            initial={{ opacity: 0, y: -10 }}
                                            animate={{ opacity: 1, y: 0 }}
                                            className="text-sm text-red-500"
                                        >
                                            {errors.user.password.message}
                                        </motion.p>
                                    )}
                                </div>
                            </motion.div>

                            {/* Submit Button */}
                            <motion.div
                                initial={{ opacity: 0, y: 20 }}
                                animate={{ opacity: 1, y: 0 }}
                                transition={{ delay: 0.5 }}
                            >
                                <Button
                                    type="submit"
                                    className="w-full bg-gradient-to-r from-violet-600 to-fuchsia-600 hover:from-violet-700 hover:to-fuchsia-700 text-white shadow-lg shadow-violet-500/50 hover:shadow-xl hover:shadow-violet-500/60 transition-all duration-300 h-12 text-base font-semibold rounded-xl"
                                    disabled={loading}
                                >
                                    {loading ? (
                                        <span className="flex items-center gap-2">
                                            <motion.div
                                                animate={{ rotate: 360 }}
                                                transition={{ duration: 1, repeat: Infinity, ease: "linear" }}
                                                className="w-5 h-5 border-2 border-white border-t-transparent rounded-full"
                                            />
                                            Creating Account...
                                        </span>
                                    ) : (
                                        <span className="flex items-center gap-2">
                                            Get Started
                                            <Rocket className="w-5 h-5" />
                                        </span>
                                    )}
                                </Button>
                            </motion.div>

                            {/* Login Link */}
                            <motion.div
                                initial={{ opacity: 0 }}
                                animate={{ opacity: 1 }}
                                transition={{ delay: 0.6 }}
                                className="text-center pt-4 border-t border-gray-200 dark:border-gray-700"
                            >
                                <p className="text-sm text-gray-600 dark:text-gray-400">
                                    Already have an account?{" "}
                                    <Link
                                        to="/login"
                                        className="text-violet-600 hover:text-violet-700 dark:text-violet-400 dark:hover:text-violet-300 font-semibold transition-colors"
                                    >
                                        Sign in
                                    </Link>
                                </p>
                            </motion.div>
                        </form>
                    </div>
                </motion.div>
            </div>

            {/* Right Side - Visual/Branding */}
            <motion.div
                initial={{ opacity: 0, x: 100 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ duration: 0.6 }}
                className="hidden lg:flex flex-1 bg-gradient-to-br from-violet-600 via-fuchsia-600 to-pink-600 p-12 items-center justify-center relative overflow-hidden"
            >
                {/* Decorative Elements */}
                <div className="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjAiIGhlaWdodD0iNjAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGRlZnM+PHBhdHRlcm4gaWQ9ImdyaWQiIHdpZHRoPSI2MCIgaGVpZ2h0PSI2MCIgcGF0dGVyblVuaXRzPSJ1c2VyU3BhY2VPblVzZSI+PHBhdGggZD0iTSAxMCAwIEwgMCAwIDAgMTAiIGZpbGw9Im5vbmUiIHN0cm9rZT0id2hpdGUiIHN0cm9rZS1vcGFjaXR5PSIwLjEiIHN0cm9rZS13aWR0aD0iMSIvPjwvcGF0dGVybj48L2RlZnM+PHJlY3Qgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNncmlkKSIvPjwvc3ZnPg==')] opacity-20" />

                <div className="relative z-10 max-w-lg">
                    <motion.div
                        initial={{ opacity: 0, y: 20 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ delay: 0.3 }}
                    >
                        <h2 className="text-5xl font-bold text-white mb-6">
                            Build the Future
                            <br />
                            <span className="text-violet-200">of Software</span>
                        </h2>
                        <p className="text-xl text-violet-100 mb-12">
                            Generate production-ready code with AI-powered workflows and intelligent automation.
                        </p>

                        {/* Features */}
                        <div className="space-y-6">
                            {features.map((feature, index) => (
                                <motion.div
                                    key={feature.title}
                                    initial={{ opacity: 0, x: -20 }}
                                    animate={{ opacity: 1, x: 0 }}
                                    transition={{ delay: 0.5 + index * 0.1 }}
                                    className="flex items-start gap-4 group cursor-pointer"
                                >
                                    <div className="w-12 h-12 rounded-xl bg-white/10 backdrop-blur-sm flex items-center justify-center group-hover:bg-white/20 transition-all duration-300 group-hover:scale-110">
                                        <feature.icon className="w-6 h-6 text-white" />
                                    </div>
                                    <div>
                                        <h3 className="text-white font-semibold text-lg mb-1">
                                            {feature.title}
                                        </h3>
                                        <p className="text-violet-100 text-sm">
                                            {feature.desc}
                                        </p>
                                    </div>
                                </motion.div>
                            ))}
                        </div>

                        {/* Floating Code Blocks Animation */}
                        <motion.div
                            className="absolute -bottom-10 -right-10 w-64 h-64"
                            animate={{
                                y: [0, -20, 0],
                            }}
                            transition={{
                                duration: 6,
                                repeat: Infinity,
                                ease: "easeInOut",
                            }}
                        >
                            <div className="w-full h-full rounded-2xl bg-white/10 backdrop-blur-md border border-white/20 p-4">
                                <div className="space-y-2">
                                    <div className="h-3 bg-white/30 rounded w-3/4" />
                                    <div className="h-3 bg-white/20 rounded w-1/2" />
                                    <div className="h-3 bg-white/30 rounded w-5/6" />
                                    <div className="h-3 bg-white/20 rounded w-2/3" />
                                </div>
                            </div>
                        </motion.div>
                    </motion.div>
                </div>
            </motion.div>
        </div>
    );
}
