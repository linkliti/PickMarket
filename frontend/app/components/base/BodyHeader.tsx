export default function BodyHeader({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <section className="mt-4 flex flex-col rounded-2xl bg-white px-4 pb-5 pt-2.5 max-md:max-w-full">
      {children}
    </section>
  );
}
