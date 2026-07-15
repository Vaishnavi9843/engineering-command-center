import MetricCard from "../../components/ui/MetricCard";

export default function DashboardPage() {
  return (
    <div>
      <h1 className="mb-8 text-3xl font-bold">
        Dashboard
      </h1>

      <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-4">
        <MetricCard title="Running Services" value="112" />
        <MetricCard title="Deployments" value="15" />
        <MetricCard title="Alerts" value="4" />
        <MetricCard title="Cloud Cost" value="$126" />
      </div>
    </div>
  );
}