const menuItems = [
  "Dashboard",
  "GitHub",
  "Kubernetes",
  "Monitoring",
  "Settings",
];

export default function Sidebar() {
  return (
    <aside className="w-64 bg-slate-900 text-white">
      <div className="border-b border-slate-700 p-6">
        <h1 className="text-xl font-bold">
          ECC
        </h1>

        <p className="mt-1 text-sm text-slate-400">
          Engineering Command Center
        </p>
      </div>

      <nav className="p-4">
        {menuItems.map((item) => (
          <button
            key={item}
            className="mb-2 w-full rounded-lg px-4 py-3 text-left transition hover:bg-slate-800"
          >
            {item}
          </button>
        ))}
      </nav>
    </aside>
  );
}