Write-Host "Testing Pet Store API" -ForegroundColor Cyan

Write-Host "`n1. Testing GET /pets..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/pets" -Method Get
    Write-Host "Success! Response:" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "Error: $_" -ForegroundColor Red
}

Write-Host "`n2. Testing POST /pets..." -ForegroundColor Yellow
$newPet = @{
    name = "Buddy"
    category = "Dog"
    price = 299.99
    status = "available"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/pets" -Method Post -Body $newPet -ContentType "application/json"
    Write-Host "Success! Response:" -ForegroundColor Green
    $response
} catch {
    Write-Host "Error: $_" -ForegroundColor Red
}

Write-Host "`nTesting completed" -ForegroundColor Cyan