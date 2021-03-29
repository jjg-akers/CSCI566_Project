## Endpoints

### /registrar

* POST
  Fields: First Name, Last Name, Some form of "ID"

  Use this for sending a new voter registration - values will be hashed and to form their unique voter ID

* GET?<voters unique ID>
  
  Use this to check if voter is registered

  ### /vote

  * POST
    send ballot selections
  * GET
    checks and validates submitted vote