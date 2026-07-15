import Card from "@/components/ui/Card";

interface MetricCardProps {
  title: string;
  value: string;
}

export default function MetricCard({
  title,
  value,
}: MetricCardProps) {
  return (
    <Card>
      <p className="text-sm text-slate-500">
        {title}
      </p>

      <h2 className="mt-2 text-3xl font-bold">
        {value}
      </h2>
    </Card>
  );
}