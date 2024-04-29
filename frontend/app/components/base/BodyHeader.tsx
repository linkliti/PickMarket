export default function BodyHeader({ children }: { children: React.ReactNode }) {
  return (
    <section className="mt-4 flex w-full grow flex-col rounded-2xl bg-white p-4 max-md:max-w-full">
      {children}
    </section>
  );
}
