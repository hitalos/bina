var bina = angular.module('bina', []);

bina.controller('buscaController', function($scope, $http){
    $http.get('/index.json').success(function(data){
        $scope.contatos = data;
    });
});
