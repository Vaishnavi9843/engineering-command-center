export default function Topbar() {
  return (
    <header className="flex h-16 items-center justify-between border-b bg-white px-8">
      <h2 className="text-xl font-semibold">
        Dashboard
      </h2>

      <div className="flex items-center gap-4">
        <button>🔔</button>

        <div className="flex h-10 w-10 items-center justify-center rounded-full bg-slate-800 text-white">
          V
        </div>
      </div>
    </header>
  );
}