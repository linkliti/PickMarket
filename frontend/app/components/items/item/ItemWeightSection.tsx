import { ItemContext } from "@/components/items/ItemContext";
import { cn } from "@/lib/utils";
import { ValueLabel, valueLabels } from "@/types/itemTypes";
import { useContext } from "react";

export default function ItemWeightSection() {
  const { maxTotalWeight, totalWeight } = useContext(ItemContext);
  return (
    <div className={cn("flex flex-col text-center")}>
      <div className="relative mt-2.5 h-[22px] shrink-0 overflow-hidden rounded-2xl border border-solid border-black ">
        <div
          className={cn(
            "bg-primary absolute h-full w-full bg-gradient-to-r from-red-500 via-yellow-300 to-green-500",
          )}
        />
        <div
          className="absolute right-0 h-full bg-white"
          style={{
            width: `${((maxTotalWeight - totalWeight) / maxTotalWeight) * 100}%`,
          }}
        />
      </div>
      <div className="text-base text-zinc-500">
        {totalWeight.toPrecision(2)}/{maxTotalWeight.toFixed(0)}
      </div>
      <div className="mt-3 self-center text-sm">
        {
          valueLabels.find(
            (label: ValueLabel): boolean =>
              label.value === Math.round((totalWeight / maxTotalWeight) * 5),
          )?.label
        }
      </div>
    </div>
  );
}
