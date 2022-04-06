import * as F from "../../build/fiber_ecommerce/main.gen";

export function App() {
  async function Test() {
    
    /** -- user test */
    console.log("GetUsers()", await (await F.GetUsers()).json());
    console.log("GetUser(1)", await (await F.GetUser(1)).json());
    console.log(
      "CreateUser (INVALID)",
      await (
        await F.CreateUser({
          username: "qweqwe",
          email: "string",
          age: 20,
          password: "string",
        })
      ).json()
    );
    console.log(
      "CreateUser (VALID)",
      await (
        await F.CreateUser({
          username: "qweqwe",
          email: "qwe@qwe.com",
          age: 20,
          password: "very-long-and-good-password",
        })
      ).json()
    );

   /** -- comment test */
    console.log("GetComment(1)", await (await F.GetComment(1)).json());
    console.log("GetComments()", await (await F.GetComments()).json());
    console.log("CreateComment()", await (await F.CreateComment()).json());
    console.log("DeleteComment()", await (await F.DeleteComment()).json());
    console.log("UpdateComment()", await (await F.UpdateComment()).json());
  }

  return (
    <>
      <div class="flex-center">
        <div onClick={Test} class="test-btn">
          Test
        </div>
      </div>
    </>
  );
}
