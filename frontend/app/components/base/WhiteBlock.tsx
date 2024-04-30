import { cn } from "@/lib/utils";
import { ReactElement } from "react";

export default function WhiteBlock({
  children,
  className = "",
}: {
  children: React.ReactNode;
  className?: string;
}): ReactElement {
  return (
    <section
      className={cn(
        "bg-card text-card-foreground mt-4 flex flex-col rounded-2xl p-4 max-md:max-w-full",
        className,
      )}
    >
      {children}
    </section>
  );
}
