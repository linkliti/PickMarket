import { Button } from "@/components/ui/button";
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { cn } from "@/lib/utils";
import { ChevronDown, ChevronRight } from "lucide-react";
import { ReactElement, useState } from "react";

export default function FilterWrapper({
  children,
  className,
  name,
  keyName,
}: {
  children: ReactElement;
  className?: string;
  name: string;
  keyName: string;
}): ReactElement {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <Collapsible
      className={cn(className)}
      open={isOpen}
      onOpenChange={setIsOpen}
      id={keyName}
    >
      <Button asChild>
        <CollapsibleTrigger
          className={cn(
            "bg-secondary flex w-full items-center justify-between gap-2 overflow-hidden",
            isOpen && "rounded-b-none",
          )}
        >
          <div className="grow truncate text-left">{name}</div>
          {/* <PrioritySelector key={filter.key} /> */}
          {isOpen ? <ChevronDown className="h-4 w-4" /> : <ChevronRight className="h-4 w-4" />}
        </CollapsibleTrigger>
      </Button>

      <CollapsibleContent className="bg-secondary">{children}</CollapsibleContent>
    </Collapsible>
  );
}
