import { StringList } from "@/proto/app/protos/types";
import { z } from "zod";

export type FilterPrefValue =
  | {
      oneofKind: "numVal";
      numVal: number;
    }
  | {
      oneofKind: "listVal";
      listVal: StringList;
    }
  | {
      oneofKind: undefined;
    };

const PrefFormSchema = z.object({
  params: z.string(),
  userQuery: z.string(),
  priorities: z.record(z.number()),
  prefs: z.record(
    z.union([z.number(), z.array(z.string()), z.boolean()]),
  ),
});

export type PrefForm = z.infer<typeof PrefFormSchema>;
