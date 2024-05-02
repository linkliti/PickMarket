import PrioritySelector from "@/components/filters/PrioritySelector";
import FilterSelectorBool from "@/components/filters/selectors/FilterSelectorBool";
import FilterSelectorList from "@/components/filters/selectors/FilterSelectorList";
import FilterSelectorRange from "@/components/filters/selectors/FilterSelectorRange";
import { Button } from "@/components/ui/button";
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { cn } from "@/lib/utils";
import { Filter } from "@/proto/app/protos/items";
import { ChevronDown, ChevronRight } from "lucide-react";
import { ReactElement, useState } from "react";

// function objectToString(obj: object): string {
//   return JSON.stringify(obj, null, 2);
// }

export default function FilterSelector({
  filter,
  className,
}: {
  filter: Filter;
  className?: string;
}): ReactElement {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <Collapsible
      className={cn(className)}
      open={isOpen}
      onOpenChange={setIsOpen}
    >
      {filter.data.oneofKind === "boolFilter" ? (
        <FilterSelectorBool filter={filter} />
      ) : (
        <Button asChild>
          <CollapsibleTrigger className="bg-secondary flex w-full items-center justify-between gap-2 overflow-hidden">
            <div className="grow truncate text-left">{filter.title}</div>
            <PrioritySelector key={filter.key} />
            {isOpen ? <ChevronDown className="h-4 w-4" /> : <ChevronRight className="h-4 w-4" />}
          </CollapsibleTrigger>
        </Button>
      )}

      <CollapsibleContent>
        {filter.data.oneofKind === "boolFilter" && <FilterSelectorBool filter={filter} />}
        {filter.data.oneofKind === "selectionFilter" && <FilterSelectorList filter={filter} />}
        {filter.data.oneofKind === "rangeFilter" && <FilterSelectorRange filter={filter} />}
      </CollapsibleContent>
    </Collapsible>
  );
}
