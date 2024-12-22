import { type PiniaPluginContext } from "pinia";

const piniaLogger = (context: PiniaPluginContext) => {
  const { store } = context;

  store.$subscribe((mutation, state) => {
    console.group(`Mutation: ${mutation.storeId}`);
    console.log("Change:", mutation);
    console.log("New State:", state);
    console.groupEnd();
  });
};

export default piniaLogger;