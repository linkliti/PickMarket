import { cn } from "@/lib/utils";
import { ReactElement } from "react";

export default function Logo({ className = "" }: { className?: string }): ReactElement {
  return <span className={cn("font-sans text-2xl font-extrabold", className)}>PickMarket</span>;
}
