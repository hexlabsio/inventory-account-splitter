import io.hexlabs.kloudformation.module.serverless.Method.*
import io.hexlabs.kloudformation.module.serverless.serverless
import io.kloudformation.KloudFormation
import io.kloudformation.StackBuilder

class Stack: StackBuilder {
    override fun KloudFormation.create(args: List<String>) {
        val codeLocation = args.first()
        serverless("account-splitter", "live", +"hexlabs-deployments") {
            globalRole { roleName("inventory-account-splitter-role") }
            serverlessFunction("account-splitter", +codeLocation,+"handler", +"go1.x"){
                http(cors = true) {
                    httpBasePathMapping(+"api.hexlabs.io",+"inventory")
                    path({ "account" / { "accountId" } }) { GET(); POST() }
                }
            }
        }
    }
}