import * as F from "../../build/fiber_ecommerce/main.gen";

export function App() {
  async function Test() {
    //
    console.log(await (await F.GetUsers()).json());
    console.log(await (await F.GetUser(1)).json());
  }

  return (
    <>
      {/* <button class="test-btn" onClick={Test}>Test</button> */}
      <div class="flex-center">
        <div onClick={Test} class="test-btn">Test</div>
      </div>
    </>
  );
}
