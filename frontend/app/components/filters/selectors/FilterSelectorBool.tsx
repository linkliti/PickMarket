import PrioritySelector from "@/components/filters/PrioritySelector";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Filter } from "@/proto/app/protos/items";
import { ReactElement } from "react";

export default function FilterSelectorBool({ filter }: { filter: Filter }): ReactElement {
  return (
    <Button
      asChild
      className="p-0"
    >
      <div className="bg-secondary flex w-full items-center justify-between overflow-hidden">
        <Label
          htmlFor={filter.key}
          className="w-10/12 truncate p-4 text-left"
        >
          {filter.title}
        </Label>
        <PrioritySelector key={filter.key} />
        <Switch
          className="mx-2 border border-gray-400"
          id={filter.key}
        />
      </div>
    </Button>
  );
}
