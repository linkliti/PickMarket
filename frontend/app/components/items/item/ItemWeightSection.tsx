import { cn } from "@/lib/utils";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { ValueLabel, valueLabels } from "@/types/itemTypes";

export default function ItemWeightSection({
  item,
  maxTotalWeight,
  className = "",
}: {
  item: ItemExtended;
  maxTotalWeight: number;
  className: string;
}) {
  return (
    <div className={cn("flex flex-col text-center", className)}>
      <div className="relative mt-2.5 h-[22px] shrink-0 overflow-hidden rounded-2xl border border-solid border-black ">
        <div
          className={cn(
            "bg-primary absolute h-full w-full bg-gradient-to-r from-red-500 via-yellow-300 to-green-500",
          )}
        />
        <div
          className="absolute right-0 h-full bg-white"
          style={{
            width: `${((maxTotalWeight - item.totalWeight) / maxTotalWeight) * 100}%`,
          }}
        />
      </div>
      <div className="text-base text-zinc-500">
        {item.totalWeight.toPrecision(3)}/{maxTotalWeight.toFixed(0)}
      </div>
      <div className="mt-3 self-center text-sm">
        {
          valueLabels.find(
            (label: ValueLabel): boolean =>
              label.value === Math.round((item.totalWeight / maxTotalWeight) * 5),
          )?.label
        }
      </div>
    </div>
  );
}
