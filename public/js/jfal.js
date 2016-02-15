var bina = angular.module('bina', []);

bina.controller('buscaController', function($scope, $http){
    $http.get('/contatos/json').success(function(data){
        $scope.contatos = data;
    });

    $('#search').focus();
});
