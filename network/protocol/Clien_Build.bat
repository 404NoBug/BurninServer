protoc.exe --csharp_out=. proto/*.proto

copy *.cs E:\myGameObject\BurninProject\Assets\Scripts

for %%i in (*.cs) do (
echo delect!path! %%i
del /a /f %%i
)

pause