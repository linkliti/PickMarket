import PrioritySelector from "@/components/filters/filterTypes/PrioritySelector";
import { Button } from "@/components/ui/button";
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { cn } from "@/lib/utils";
import { PrefForm } from "@/types/filterTypes";
import { ChevronDown, ChevronRight } from "lucide-react";
import { ReactElement, useState } from "react";
import { Control } from "react-hook-form";

export default function FilterWrapper({
  children,
  className,
  name,
  keyName,
  control,
}: {
  children: ReactElement;
  className?: string;
  name: string;
  keyName: string;
  control: Control<PrefForm, unknown>;
}): ReactElement {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <Collapsible
      className={cn(
        isOpen ? "" : "rounded-b-none",
        "rounded-sm border-b border-b-gray-300",
        className,
      )}
      open={isOpen}
      onOpenChange={setIsOpen}
      id={keyName}
    >
      <Button
        asChild
        className={cn(
          "bg-secondary flex w-full items-center justify-between gap-2 overflow-hidden p-0",
          isOpen && "rounded-b-none",
        )}
      >
        <div className="pr-4">
          <CollapsibleTrigger
            className={cn("flex w-full items-center justify-between gap-2 overflow-hidden")}
          >
            <>
              <p className=" grow truncate py-1 pl-4 text-left leading-8">{name}</p>
              {isOpen ? <ChevronDown className="h-4 w-4" /> : <ChevronRight className="h-4 w-4" />}
            </>
          </CollapsibleTrigger>
          <PrioritySelector
            keyName={keyName}
            control={control}
          />
        </div>
      </Button>
      <CollapsibleContent className="bg-secondary">{children}</CollapsibleContent>
    </Collapsible>
  );
}
