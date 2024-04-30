import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { marketplaces } from "@/store/categoryStore";
import { Check, ChevronsUpDown } from "lucide-react";
import { ReactElement, useState } from "react";

interface Marketplace {
  label: string;
  value: string;
}

export default function CategoryMarketSelector({
  className = "",
  listClassNames = "",
}: {
  className?: string;
  listClassNames?: string;
}): ReactElement {
  const [isOpenMarketSelect, setOpenMarketSelect] = useState(false);
  const [selectedMarket, setSelectedMarket] = useState(marketplaces[0]);

  return (
    <div className={cn(className)}>
      <Popover
        open={isOpenMarketSelect}
        onOpenChange={setOpenMarketSelect}
      >
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            disabled={marketplaces.length < 2}
            role="combobox"
            aria-expanded={isOpenMarketSelect}
            className={cn("justify-between", listClassNames)}
          >
            {selectedMarket
              ? marketplaces.find(
                  (market: Marketplace): boolean => market.value === selectedMarket.value,
                )?.label
              : "Выбрать маркетплейс..."}
            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className={cn("p-0", listClassNames)}>
          <Command>
            {/* <CommandInput placeholder="Выбрать маркетплейс..." /> */}
            <CommandList>
              <CommandEmpty>Маркетплейс не найден</CommandEmpty>
              <CommandGroup>
                {marketplaces.map(
                  (market: Marketplace): ReactElement => (
                    <CommandItem
                      key={market.value}
                      value={market.value}
                      onSelect={(currentValue: string): void => {
                        setSelectedMarket(
                          marketplaces.find(
                            (market: Marketplace): boolean => market.value === currentValue,
                          )!,
                        );
                        setOpenMarketSelect(false);
                      }}
                    >
                      <Check
                        className={cn(
                          "mr-2 h-4 w-4",
                          selectedMarket.value === market.value ? "opacity-100" : "opacity-0",
                        )}
                      />
                      {market.label}
                    </CommandItem>
                  ),
                )}
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
    </div>
  );
}
