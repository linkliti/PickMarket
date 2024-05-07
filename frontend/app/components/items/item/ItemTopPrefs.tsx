import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { Characteristic } from "@/proto/app/protos/items";
import { MinusSquareIcon, PlusSquareIcon } from "lucide-react";
import { ReactElement, useState } from "react";

function printChar(char: Characteristic) {
  if (char.value.oneofKind === "listVal") {
    const listValStr = char.value.listVal.values.map((value) => value).join(", ");
    if (listValStr) {
      return listValStr;
    }
  } else if (char.value.oneofKind === "numVal") {
    return char.value.numVal.toString();
  }
  return "";
}

export default function ItemTopPrefs({ chars }: { chars: Characteristic[] }): ReactElement {
  const weightedChars = chars.filter((char) => char.charWeight !== 0);
  const nonWeightedChars = chars.filter((char) => char.charWeight === 0);

  const [isOpen, setIsOpen] = useState(false);
  return (
    <>
      <table className="w-full table-auto">
        <thead></thead>
        <tbody>
          <TableChars
            chars={weightedChars}
            plusIcon={<PlusSquareIcon className="size-4 fill-green-400" />}
          />
          <Collapsible
            open={isOpen}
            onOpenChange={setIsOpen}
            className=""
          >
            <CollapsibleTrigger>
              <p className="font-bold hover:underline">{isOpen ? "Скрыть" : "Показать все"}</p>
            </CollapsibleTrigger>
            <CollapsibleContent>
              <TableChars
                chars={nonWeightedChars}
                plusIcon={<MinusSquareIcon className="size-4 " />}
              />
            </CollapsibleContent>
          </Collapsible>
        </tbody>
      </table>
    </>
  );
}

function TableChars({
  chars,
  plusIcon: iconElement,
}: {
  chars: Characteristic[];
  plusIcon: ReactElement;
}): ReactElement {
  return (
    <>
      {chars.map(
        (char: Characteristic): ReactElement => (
          <tr
            key={char.key}
            className="flex flex-row gap-2 border-b border-b-gray-200"
          >
            <td className="flex w-3/5 items-baseline">
              {iconElement}
              <p className="w-1/2 grow text-wrap ps-2">{char.name + ": "}</p>
            </td>
            <td className="flex w-2/5 items-center gap-2 text-wrap">
              <span className="r-0">{printChar(char)}</span>
            </td>
          </tr>
        ),
      )}
    </>
  );
}
